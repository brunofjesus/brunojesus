package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/a-h/templ"
	chromahtml "github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/styles"
	"github.com/brunofjesus/brunojesus/internal/layout"
	highlighting "github.com/yuin/goldmark-highlighting/v3"
	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
	"golang.org/x/sync/errgroup"
)

var layoutSettings = layout.Settings{
	SiteName: "Bruno Jesus",
	SubTitle: "Yet another boring minimalist developer",
	HeaderLinks: []layout.Link{
		{Name: "Home", Href: "/"},
		{Name: "Blog", Href: "/articles"},
		{Name: "CV", Href: "/cv"},
		{Name: "Links", Href: "/links"},
	},
}

const (
	MetaTitleKey    = "Title"
	MetaDateKey     = "Date"
	MetaFileNameKey = "FileName"
)

var (
	ErrNoTitleMetadata    = errors.New("no title metadata")
	ErrNoDateMetadata     = errors.New("no date metadata")
	ErrInvalidDateFormat  = errors.New("invalid date format")
	ErrNoFileNameMetadata = errors.New("no file name metadata")
)

var markdownParser = parser.New(
	parser.WithExtensions(
		highlighting.Parser,
		meta.Parser,
		extension.NewTableParser(),
	),
)

var htmlRenderer = html.New(
	html.WithUnsafe(),
	html.WithExtensions(
		highlighting.NewHTMLRenderer(
			highlighting.WithFormatterOptions(
				chromahtml.WithClasses(true),
				chromahtml.WithLineNumbers(true),
			),
		),
		extension.NewTableHTMLRenderer(),
	),
)

var rssRenderer = html.New(
	html.WithUnsafe(),
	html.WithExtensions(
		extension.NewTableHTMLRenderer(),
	),
)

func main() {
	var g errgroup.Group
	g.Go(generateStylesheets)
	g.Go(renderPages)

	if err := g.Wait(); err != nil {
		log.Fatalf("error generating site: %v", err)
	}
}

// markdownFile holds the paths derived from a source markdown file, computed
// once so downstream code never has to re-derive them from strings.
type markdownFile struct {
	sourcePath   string // e.g. data/articles/foo.md
	relativePath string // e.g. articles/foo.md
	slug         string // e.g. articles/foo
	outputDir    string // e.g. gen/articles
	outputPath   string // e.g. gen/articles/foo.html
	isArticle    bool
}

func newMarkdownFile(path string) (markdownFile, error) {
	relativePath, err := filepath.Rel("data", path)
	if err != nil {
		return markdownFile{}, fmt.Errorf("failed to get relative path for %s: %w", path, err)
	}

	relativeDir := filepath.Dir(relativePath)
	htmlFileName := strings.TrimSuffix(filepath.Base(relativePath), ".md") + ".html"

	return markdownFile{
		sourcePath:   path,
		relativePath: relativePath,
		slug:         strings.TrimSuffix(relativePath, ".md"),
		outputDir:    filepath.Join("gen", relativeDir),
		outputPath:   filepath.Join("gen", relativeDir, htmlFileName),
		isArticle:    relativeDir == "articles",
	}, nil
}

func renderPages() error {
	var metadataList []layout.PageMetadata
	var rssItems []templ.Component

	err := filepath.WalkDir("data", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		mf, err := newMarkdownFile(path)
		if err != nil {
			return err
		}

		metadata, rssItem, err := renderMarkdownFile(mf)
		if err != nil {
			return err
		}
		if metadata != nil {
			metadataList = append(metadataList, *metadata)
		}
		if rssItem != nil {
			rssItems = append(rssItems, rssItem)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through data directory: %w", err)
	}

	if err := renderArticlePage(metadataList); err != nil {
		return fmt.Errorf("error rendering articles page: %w", err)
	}

	if err := renderRSSFeed(rssItems); err != nil {
		return fmt.Errorf("error rendering RSS feed: %w", err)
	}

	return nil
}

func renderMarkdownFile(mf markdownFile) (*layout.PageMetadata, templ.Component, error) {
	source, err := os.ReadFile(mf.sourcePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file %s: %w", mf.sourcePath, err)
	}

	doc := markdownParser.Parse(source)

	var metadata layout.PageMetadata
	if mf.isArticle {
		metadata, err = metadataFromDocument(doc, mf.relativePath)
		if err != nil {
			return nil, nil, err
		}
	}

	if err := os.MkdirAll(mf.outputDir, os.ModePerm); err != nil {
		return nil, nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	var htmlBuf bytes.Buffer
	if err := htmlRenderer.Render(&htmlBuf, source, doc); err != nil {
		return nil, nil, fmt.Errorf("failed to render HTML for file %s: %w", mf.sourcePath, err)
	}

	f, err := os.Create(mf.outputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create HTML file for %s: %w", mf.sourcePath, err)
	}
	defer f.Close()

	err = layout.Base(
		layoutSettings,
		layout.Page(metadata.Title, htmlBuf.String(), metadata.Date),
	).Render(context.Background(), f)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to render layout for file %s: %w", mf.sourcePath, err)
	}

	if !mf.isArticle {
		return nil, nil, nil
	}

	var rssBuf bytes.Buffer
	if err := rssRenderer.Render(&rssBuf, source, doc); err != nil {
		return nil, nil, fmt.Errorf("failed to render RSS HTML for file %s: %w", mf.sourcePath, err)
	}
	rssItem := layout.RSSItem(
		metadata.Title, mf.slug, rssBuf.String(), *metadata.Date,
	)

	return &metadata, rssItem, nil
}

func renderArticlePage(metadata []layout.PageMetadata) error {
	f, err := os.Create("gen/articles.html")
	if err != nil {
		return fmt.Errorf("failed to create articles.html: %w", err)
	}
	defer f.Close()

	err = layout.Base(
		layoutSettings,
		layout.Articles(metadata),
	).Render(context.Background(), f)
	if err != nil {
		return fmt.Errorf("failed to render articles.html: %w", err)
	}

	return nil
}

func renderRSSFeed(rssItems []templ.Component) error {
	f, err := os.Create("gen/rss.xml")
	if err != nil {
		return fmt.Errorf("failed to create rss.xml: %w", err)
	}
	defer f.Close()

	err = layout.RSSFeed(layoutSettings, rssItems).Render(context.Background(), f)
	if err != nil {
		return fmt.Errorf("failed to render rss.xml: %w", err)
	}

	return nil
}

func metadataFromDocument(doc ast.Node, fileName string) (layout.PageMetadata, error) {
	metadata := doc.(*ast.Document).Metadata()
	title, ok := metadata[MetaTitleKey].(string)
	if !ok {
		return layout.PageMetadata{}, ErrNoTitleMetadata
	}

	dateString, ok := metadata[MetaDateKey].(string)
	if !ok {
		return layout.PageMetadata{}, ErrNoDateMetadata
	}

	date, err := time.Parse(time.RFC1123Z, dateString)
	if err != nil {
		return layout.PageMetadata{}, ErrInvalidDateFormat
	}

	return layout.PageMetadata{
		Title:    title,
		Date:     &date,
		FileName: strings.TrimSuffix(fileName, ".md"),
	}, nil
}

func generateStylesheets() error {
	themes := []string{
		"catppuccin-frappe",
		"catppuccin-latte",
	}

	for _, theme := range themes {
		style := styles.Get(theme)
		formatter := chromahtml.New(
			chromahtml.WithClasses(true),
			chromahtml.WithLineNumbers(true),
		)

		f, err := os.Create("assets/css/code-" + theme + ".css")
		if err != nil {
			return err
		}
		defer f.Close()

		if err := formatter.WriteCSS(f, style); err != nil {
			return err
		}
	}

	return nil
}

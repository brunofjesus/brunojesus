package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	chromahtml "github.com/alecthomas/chroma/v3/formatters/html"
	"github.com/alecthomas/chroma/v3/styles"
	"github.com/brunofjesus/brunojesus/internal/layout"
	highlighting "github.com/yuin/goldmark-highlighting/v3"
	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

var (
	ErrNoTitleMetadata    = errors.New("no title metadata")
	ErrNoDateMetadata     = errors.New("no date metadata")
	ErrInvalidDateFormat  = errors.New("invalid date format")
	ErrNoFileNameMetadata = errors.New("no file name metadata")
)

const (
	MetaTitleKey    = "Title"
	MetaDateKey     = "Date"
	MetaFileNameKey = "FileName"
)

func main() {
	err := generateStylesheets()
	if err != nil {
		fmt.Println(err)
	}

	err = renderPages()
	if err != nil {
		fmt.Println(err)
	}
}

func renderPages() error {
	var metadataList []layout.PageMetadata

	markdownParser := parser.New(
		parser.WithExtensions(
			highlighting.Parser,
			meta.Parser,
		),
	)

	htmlRenderer := html.New(
		html.WithUnsafe(),
		html.WithExtensions(
			highlighting.NewHTMLRenderer(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatterOptions(
					chromahtml.WithClasses(true),
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	err := filepath.WalkDir("data", func(path string, d fs.DirEntry, werr error) error {
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		source, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		doc := markdownParser.Parse(source)

		relativePath, err := filepath.Rel("data", path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		var metadata layout.PageMetadata
		if filepath.Dir(relativePath) == "articles" {
			metadata, err = metadataFromDocument(doc, relativePath)
			if err != nil {
				return err
			}
			metadataList = append(metadataList, metadata)
		}

		targetDir := filepath.Join("gen", filepath.Dir(relativePath))
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		var buf bytes.Buffer
		if err := htmlRenderer.Render(&buf, source, doc); err != nil {
			return fmt.Errorf("failed to render HTML for file %s: %w", path, err)
		}

		f, err := os.Create("gen/" + strings.TrimSuffix(relativePath, ".md") + ".html")
		if err != nil {
			return fmt.Errorf("failed to create HTML file for %s: %w", path, err)
		}
		defer f.Close()

		err = layout.Page(metadata.Title, buf.String(), metadata.Date).Render(context.Background(), f)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through data directory: %w", err)
	}

	// Generate the list
	f, err := os.Create("gen/articles.html")
	if err != nil {
		return fmt.Errorf("failed to create articles.html: %w", err)
	}
	defer f.Close()
	err = layout.Articles(metadataList).Render(context.Background(), f)
	if err != nil {
		return fmt.Errorf("failed to render articles.html: %w", err)
	}

	return err
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
		"monokai",
		"gruvbox",
		"gruvbox-light",
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

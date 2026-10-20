package pages

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

//go:embed *.md articles/*.md
var fs embed.FS

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

type PageMetadata struct {
	Title    string
	Date     *time.Time
	FileName string
}

type Page struct {
	PageMetadata
	Content string
}

func NewPageMetadata(metadata map[string]any) (PageMetadata, error) {
	title, ok := metadata[MetaTitleKey].(string)
	if !ok {
		return PageMetadata{}, ErrNoTitleMetadata
	}

	dateString, ok := metadata[MetaDateKey].(string)
	if !ok {
		return PageMetadata{}, ErrNoDateMetadata
	}

	date, err := time.Parse(time.RFC1123Z, dateString)
	if err != nil {
		return PageMetadata{}, ErrInvalidDateFormat
	}

	fileName, ok := metadata[MetaFileNameKey].(string)
	if !ok {
		return PageMetadata{}, ErrNoFileNameMetadata
	}

	return PageMetadata{
		Title:    title,
		Date:     &date,
		FileName: fileName,
	}, nil
}

func (p PageMetadata) RenderHTMLTitle(b *bytes.Buffer) error {
	if len(p.Title) == 0 {
		return ErrNoTitleMetadata
	}

	fmt.Fprintf(b, "<h1><a href=\"#\">%s</a>", p.Title)
	if p.Date != nil {
		fmt.Fprintf(b, "%s", p.Date.Format(" — 02 January, 2006"))
	}
	fmt.Fprintf(b, "</h1>")
	return nil
}

func getDocument(fileName string) ([]byte, ast.Node, error) {
	p := parser.New(
		parser.WithExtensions(
			meta.Parser,
		),
	)

	source, err := fs.ReadFile(strings.ToLower(fileName))
	if err != nil {
		return nil, nil, err
	}

	return source, p.Parse(source), nil
}

func RenderPage(fileName string) (Page, error) {
	r := html.New(
		html.WithUnsafe(),
		html.WithExtensions(),
	)

	source, doc, err := getDocument(fileName)
	if err != nil {
		return Page{}, err
	}

	var content bytes.Buffer
	rawMeta := doc.(*ast.Document).Metadata()
	rawMeta[MetaFileNameKey] = fileName
	meta, err := NewPageMetadata(rawMeta)
	if err != nil {
		slog.Warn("Failed to get metadata for page", "fileName", fileName, "error", err)
	}

	if err := meta.RenderHTMLTitle(&content); err != nil {
		slog.Warn("Failed to render HTML title for page", "fileName", fileName, "error", err)
	}

	if err := r.Render(&content, source, doc); err != nil {
		return Page{}, err
	}

	return Page{
		PageMetadata: meta,
		Content:      content.String(),
	}, nil
}

func ListArticles() ([]PageMetadata, error) {
	entries, err := fs.ReadDir("articles")
	if err != nil {
		return nil, err
	}

	result := make([]PageMetadata, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".md") {
			_, doc, err := getDocument("articles/" + entry.Name())
			if err != nil {
				slog.Warn("Failed to get document for page", "fileName", entry.Name(), "error", err)
				continue
			}
			rawMeta := doc.(*ast.Document).Metadata()
			rawMeta[MetaFileNameKey] = strings.TrimSuffix(entry.Name(), ".md")
			meta, err := NewPageMetadata(rawMeta)
			if err != nil {
				slog.Warn("Failed to get metadata for page", "fileName", entry.Name(), "error", err)
				continue
			}

			result = append(result, meta)
		}
	}

	return result, nil
}

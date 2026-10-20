package engine

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/brunofjesus/brunojesus/internal/pages"
	meta "github.com/yuin/goldmark-meta/v2"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

type Settings struct {
	SiteName    string
	SubTitle    string
	HeaderLinks []link
}

var defaultSettings Settings = Settings{
	SiteName: "Bruno Jesus",
	SubTitle: "Yet another boring minimalist developer",
	HeaderLinks: []link{
		{"Home", "/"},
		{"Blog", "/blog"},
		{"CV", "/cv"},
		{"Links", "/links"},
	},
}

type link struct {
	name string
	link string
}

// RenderPage renders the markdown content of internal/pages/<page>.md into
// the site layout.
func RenderPage(title string) (templ.Component, error) {
	p := parser.New(
		parser.WithExtensions(
			meta.Parser,
		),
	)

	r := html.New(
		html.WithUnsafe(),
		html.WithExtensions(),
	)

	source, err := pages.FS.ReadFile(strings.ToLower(title) + ".md")
	if err != nil {
		return nil, err
	}

	doc := p.Parse(source)

	var content bytes.Buffer
	header, err := getArticleTitle(doc.(*ast.Document).Metadata())
	if err != nil {
		fmt.Println("Warning: ", err)
	} else {
		content.WriteString(header)
	}

	if err := r.Render(&content, source, doc); err != nil {
		return nil, err
	}

	return layout(defaultSettings, title, content.String()), nil
}

func getArticleTitle(metadata map[string]any) (string, error) {
	result, ok := metadata[pages.MetaTitleKey].(string)
	if !ok {
		return "", errors.New("no title metadata")
	}

	result = fmt.Sprintf("<a href=\"#\">%s</a>", result)

	dateString, ok := metadata[pages.MetaDateKey].(string)
	if !ok {
		return "", errors.New("no date metadata")
	}

	date, err := time.Parse(time.RFC1123Z, dateString)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}

	result += " — " + date.Format("02 January, 2006")

	return fmt.Sprintf("<h1>%s</h1>", result), nil
}

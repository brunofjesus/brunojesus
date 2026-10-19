package engine

import (
	"bytes"
	"strings"

	"github.com/a-h/templ"
	"github.com/brunofjesus/brunojesus/internal/pages"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
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

// markdown renders page content to HTML, allowing raw HTML blocks (such as
// the timeline and skills sections) to pass through untouched.
var markdown = goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))

// RenderPage renders the markdown content of internal/pages/<page>.md into
// the site layout.
func RenderPage(title string) (templ.Component, error) {
	source, err := pages.FS.ReadFile(strings.ToLower(title) + ".md")
	if err != nil {
		return nil, err
	}

	var content bytes.Buffer
	if err := markdown.Convert(source, &content); err != nil {
		return nil, err
	}

	return layout(defaultSettings, title, content.String()), nil
}

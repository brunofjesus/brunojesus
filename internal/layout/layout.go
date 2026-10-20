package layout

import (
	"github.com/a-h/templ"
	"github.com/brunofjesus/brunojesus/internal/pages"
)

type Settings struct {
	SiteName    string
	SubTitle    string
	HeaderLinks []Link
}

type Link struct {
	Name string
	Href string
}

func Render(settings Settings, page pages.Page) templ.Component {
	return layout(settings, page.Title, page.Content)
}

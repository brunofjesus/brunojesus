package engine

import (
	"github.com/a-h/templ"
)

type PageRenderRequest struct {
	Title string
}

func RenderPage(title string) templ.Component {
	return layout(title)
}

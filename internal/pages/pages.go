package pages

import "embed"

//go:embed *.md articles/*.md
var FS embed.FS

const (
	MetaTitleKey = "Title"
	MetaDateKey  = "Date"
)

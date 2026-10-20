// Package static embeds the generated site assets and pages. It lives at
// the repository root (next to go.mod) because go:embed patterns resolve
// relative to the source file's own directory, and that's where assets/
// and gen/ live. The go:generate directives live here for the same reason:
// "go generate" runs in the directory of the file containing the directive.
package static

//go:generate go tool templ generate
//go:generate go run ./cmd/generator

import "embed"

//go:embed assets/*
var Assets embed.FS

//go:embed gen/*
var Gen embed.FS

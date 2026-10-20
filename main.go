//go:generate go tool templ generate
//go:generate go run cmd/generator/main.go
package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"path/filepath"

	"github.com/brunofjesus/brunojesus/internal/layout"
)

//go:embed assets/*
var assets embed.FS

//go:embed gen/*
var gen embed.FS

var layoutSettings layout.Settings = layout.Settings{
	SiteName: "Bruno Jesus",
	SubTitle: "Yet another boring minimalist developer",
	HeaderLinks: []layout.Link{
		{Name: "Home", Href: "/"},
		{Name: "Blog", Href: "/articles"},
		{Name: "CV", Href: "/cv"},
		{Name: "Links", Href: "/links"},
	},
}

func servePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("page")
		if len(name) == 0 {
			name = "home"
		}
		name = name + ".html"

		path := filepath.Join("gen", name)
		f, err := gen.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		component := layout.Layout(layoutSettings, "", string(f))
		if err := component.Render(context.Background(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{page...}", servePage())
	mux.Handle("/assets/", http.FileServer(http.FS(assets)))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"bytes"
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/brunofjesus/brunojesus/internal/layout"
	"github.com/brunofjesus/brunojesus/internal/pages"
)

//go:embed assets/*
var content embed.FS

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
		name = name + ".md"

		page, err := pages.RenderPage(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		component := layout.Render(layoutSettings, page)
		if err := component.Render(context.Background(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		}
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{page...}", servePage())

	mux.Handle("/assets/", http.FileServer(http.FS(content)))

	mux.HandleFunc("/articles/{$}", func(w http.ResponseWriter, r *http.Request) {
		articleMetadataList, err := pages.ListArticles()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		articleListComponent := layout.Articles(articleMetadataList)

		var articleList bytes.Buffer
		if err := articleListComponent.Render(context.Background(), &articleList); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		component := layout.Render(layoutSettings, pages.Page{
			Title:   "Articles",
			Content: articleList.String(),
		})

		if err := component.Render(context.Background(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

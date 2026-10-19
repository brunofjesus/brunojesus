package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/brunofjesus/brunojesus/internal/engine"
)

//go:embed assets/*
var content embed.FS

func servePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := r.PathValue("page")
		if len(page) == 0 {
			page = "home"
		}

		component, err := engine.RenderPage(page)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := component.Render(context.Background(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServer(http.FS(content)))

	mux.HandleFunc("/{$}", servePage())
	mux.HandleFunc("/{page}", servePage())

	log.Fatal(http.ListenAndServe(":8080", mux))
}

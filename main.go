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

func main() {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServer(http.FS(content)))

	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		err := engine.RenderPage("yo").Render(context.Background(), w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

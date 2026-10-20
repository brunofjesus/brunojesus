package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"

	static "github.com/brunofjesus/brunojesus"
)

func servePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("page")
		if len(name) == 0 {
			name = "index.html"
		} else if filepath.Ext(name) == "" {
			name = name + ".html"
		}

		path := filepath.Join("gen", name)
		f, err := static.Gen.ReadFile(path)
		if err != nil && errors.Is(err, fs.ErrNotExist) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("404 page not found"))
			slog.Warn("Page not found", "path", path)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unexpected error"))
			slog.Error("Unexpected error reading file", "path", path, "error", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(f)
	}
}

func main() {
	helpFlag := flag.Bool("help", false, "Show help")
	listenAddrFlag := flag.String("listen", ":8080", "Address to listen on")
	flag.Parse()

	if *helpFlag {
		fmt.Println("Usage: server [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Starting server on %s\n", *listenAddrFlag)
	mux := http.NewServeMux()
	mux.HandleFunc("/{page...}", servePage())
	mux.Handle("/assets/", http.FileServer(http.FS(static.Assets)))
	log.Fatal(http.ListenAndServe(*listenAddrFlag, mux))
}

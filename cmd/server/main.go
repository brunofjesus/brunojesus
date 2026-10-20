package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	static "github.com/brunofjesus/brunojesus"
	"github.com/brunofjesus/brunojesus/internal/layout"
)

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
		f, err := static.Gen.ReadFile(path)
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

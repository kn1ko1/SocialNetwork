package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"socialnetwork/api"
	"socialnetwork/repo"
	"socialnetwork/router"
	"socialnetwork/ui"
)

const (
	addr = ":8080"
)

func main() {
	// Setup serve mux
	mux := http.NewServeMux()
	// Host static files
	serveStaticFiles(mux)
	// Add handlers to router
	setupRouter(mux)
	// Listen and serve
	fmt.Printf("server listening at address %s...\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setupRouter(mux *http.ServeMux) {
	rt := router.NewRouter()
	// Dummy UI handler, as an example
	rt.AddHandler(regexp.MustCompile(`^/$`), ui.NewDummyPageHandler())
	r := repo.NewSQLiteRepository()
	commentsHandler := api.NewCommentsHandler(r)
	mux.Handle("/", rt)
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

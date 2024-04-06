package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(a.LoadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	// register routes
	mux.Get("/", a.homeHandler)

	mux.Handle("/public/*", http.StripPrefix("/public/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveFiles(w, r, "./public/")
	})))

	return mux
}

// Serve files with proper MIME types
func serveFiles(w http.ResponseWriter, r *http.Request, dir string) {
	path := r.URL.Path

	// Determine the MIME type based on file extension
	var contentType string
	switch {
	case filepath.Ext(path) == ".css":
		contentType = "text/css"
	default:
		contentType = "text/plain"
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", contentType)

	// Serve the file using ServeFile
	http.ServeFile(w, r, filepath.Join(dir, path))
}

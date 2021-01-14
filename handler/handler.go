package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mikesparr/ai-demo-ingest/message"
	"net/http"
)

var producer message.Producer

// NewHandler instantiates a handler and injects a message producer
func NewHandler(p message.Producer) http.Handler {
	router := chi.NewRouter()
	producer = p

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/notes", batch)
	return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	err := render.Render(w, r, ErrMethodNotAllowed)
	if err != nil {
		fmt.Println("Error rendering")
	}
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	err := render.Render(w, r, ErrNotFound)
	if err != nil {
		fmt.Println("Error rendering")
	}
}

package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mikesparr/ai-demo-ingest/message"
	"net/http"
)

var producer message.Producer

func NewHandler(p message.Producer) http.Handler {
	router := chi.NewRouter()
	producer = p
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/notes", batch)
	return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}

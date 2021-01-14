package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mikesparr/ai-demo-ingest/models"
	"net/http"
)

func batch(router chi.Router) {
	router.Post("/", submitBatch)
}

func submitBatch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Submitting batch")
	batch := &models.Batch{}
	if err := render.Bind(r, batch); err != nil {
		err := render.Render(w, r, ErrorBadRequest(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := producer.SubmitBatch(batch); err != nil {
		err := render.Render(w, r, ErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
	if err := render.Render(w, r, batch); err != nil {
		err := render.Render(w, r, ServerErrorRenderer(err))
		if err != nil {
			fmt.Println("Error rendering")
		}
		return
	}
}

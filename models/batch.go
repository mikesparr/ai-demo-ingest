package models

import (
	"fmt"
	"net/http"
)

type Note struct {
	ID       string    `json:"id"`             // uuid
	Name     string    `json:"name,omitempty"` // uuid
	Features []float64 `json:"features"`
}
type Batch struct {
	ID    string `json:"id,omitempty"`
	Notes []Note `json:"notes"`
}

func (b *Batch) Bind(r *http.Request) error {
	if b.Notes == nil {
		return fmt.Errorf("notes is a required field")
	}
	return nil
}
func (*Batch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

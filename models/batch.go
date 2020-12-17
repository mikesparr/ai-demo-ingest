package models

import (
	"fmt"
	"net/http"
)

type Note struct {
	ID        string    `json:"id"`         // uuid
	SubjectID string    `json:"subject_id"` // uuid
	Features  []float64 `json:"features"`
	CreatedAt string    `json:"created_at"`
}
type Batch struct {
	ID        string `json:"id"`
	Notes     []Note `json:"notes"`
	CreatedAt string `json:"created_at"`
}
type BatchList struct {
	Batches []Batch `json:"batches"`
}

func (b *Batch) Bind(r *http.Request) error {
	if b.Notes == nil {
		return fmt.Errorf("notes is a required field")
	}
	return nil
}
func (*BatchList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Batch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

package message

import (
	"fmt"
	"github.com/mikesparr/ai-demo-ingest/models"
)

func (producer Producer) SubmitBatch(batch *models.Batch) error {
	var id string
	var createdAt string
	fmt.Println("I ran SubmitBatch !!!")
	batch.ID = id
	batch.CreatedAt = createdAt
	return nil
}

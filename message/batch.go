package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mikesparr/ai-demo-ingest/models"

	"cloud.google.com/go/pubsub"
)

// SubmitBatch publishes batch  to pubsub topic
func (producer Producer) SubmitBatch(batch *models.Batch) error {
	ctx := context.Background()

	// add new uuid for "request_id" before sending to pubsub
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	batch.ID = uuidV1.String()
	fmt.Printf("New batch submission %s\n", uuidV1.String())

	topic := producer.Topic
	batchJSON, err := json.Marshal(batch)
	if err != nil {
		return err
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(batchJSON),
	})
	if _, err := res.Get(ctx); err != nil {
		return err
	}

	return nil
}

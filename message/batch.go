package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mikesparr/ai-demo-ingest/models"

	"cloud.google.com/go/pubsub"
)

func (producer Producer) SubmitBatch(batch *models.Batch) error {
	ctx := context.Background()

	// add new uuid for "request_id" before sending to pubsub
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	batch.RequestID = uuidV1.String()
	fmt.Printf("New batch submission %s\n", uuidV1.String())

	topic := producer.Topic
	batchJson, err := json.Marshal(batch)
	if err != nil {
		return err
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(batchJson),
	})
	if _, err := res.Get(ctx); err != nil {
		return err
	}

	return nil
}

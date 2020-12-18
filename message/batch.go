package message

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mikesparr/ai-demo-ingest/models"

	"cloud.google.com/go/pubsub"
)

func (producer Producer) SubmitBatch(batch *models.Batch) error {
	var id string
	var createdAt string
	ctx := context.Background()
	fmt.Println("I ran SubmitBatch !!!")

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

	batch.ID = id
	batch.CreatedAt = createdAt
	return nil
}

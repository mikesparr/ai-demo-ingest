package message

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Producer provides message topic to inject
type Producer struct {
	Topic *pubsub.Topic
}

// Initialize connects to pubsub
func Initialize(projectID, topicID string) (Producer, error) {
	producer := Producer{}
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return producer, err
	}
	t := client.Topic(topicID)

	producer.Topic = t

	fmt.Println("Pubsub topic initialized")
	return producer, nil
}

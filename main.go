package main

import (
	"context"
	"fmt"
	"github.com/mikesparr/ai-demo-ingest/handler"
	"github.com/mikesparr/ai-demo-ingest/message"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// server
	addr := ":8080"
	/* #nosec */
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	// pubsub
	projectID, topicID :=
		os.Getenv("PROJECT_ID"),
		os.Getenv("TOPIC_ID")
	producer, err := message.Initialize(projectID, topicID)
	if err != nil {
		log.Fatalf("Could not set up messaging: %v", err)
	}
	defer producer.Topic.Stop()

	// inject pubsub producer
	httpHandler := handler.NewHandler(producer)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		err := server.Serve(listener)
		if err != nil {
			log.Printf("Error starting the server %v\n", err)
		}
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

// Stop safely shuts down server
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}

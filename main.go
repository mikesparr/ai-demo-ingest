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
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	projectID, topicID :=
		os.Getenv("PROJECT_ID"),
		os.Getenv("TOPIC_ID")
	producer, err := message.Initialize(projectID, topicID)
	if err != nil {
		log.Fatalf("Could not set up messaging: %v", err)
	}
	defer producer.Topic.Stop()

	httpHandler := handler.NewHandler(producer)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}

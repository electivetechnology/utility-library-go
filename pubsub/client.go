package pubsub

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	MAX_MESSAGES = 10 // Maximum number of messages to book at a time
	SYNCHRONOUS  = false
)

var log logger.Logging

type Client struct {
	projectID    string
	PubSubClient *pubsub.Client
	Ctx          context.Context
	Synchronous  bool
	MaxMessages  int
}

func init() {
	// Add generic logger
	log = logger.NewLogger("pubsub")
}

func NewClient(ctx context.Context) (c *Client, err error) {
	// Get Project id
	log.Printf("Configuring new PubSub client")
	projectID := os.Getenv("GOOGLE_PROJECT_ID")

	// Create  new client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("pubsub.NewClient: %v", err)
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}

	log.Printf("PubSub client configured successfully")

	return &Client{projectID, client, ctx, SYNCHRONOUS, MAX_MESSAGES}, nil
}

func NewAsyncClient() (*Client, error) {
	// Subscribe to PubSub channel
	ctx := context.Background()
	client, err := NewClient(ctx)
	client.Synchronous = true
	if err != nil {
		log.Printf("Failed to Initialize Pub/Sub Client: %v", err)
		return client, err
	}

	return client, nil
}

func (psClient *Client) SetContext(ctx context.Context) {
	psClient.Ctx = ctx
}

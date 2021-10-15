package bigquery

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
)

type Client struct {
	projectID string
	BQClient  *bigquery.Client
	ctx       context.Context
}

func NewClient(ctx context.Context) (c *Client, err error) {
	log.Printf("Configuring new BigQuery client")
	projectID := os.Getenv("GOOGLE_PROJECT_ID")

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return &Client{}, fmt.Errorf("bigquery.NewClient: %v", err)
	}

	defer client.Close()

	return &Client{projectID, client, ctx}, nil
}

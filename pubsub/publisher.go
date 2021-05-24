package pubsub

import (
	"fmt"

	"cloud.google.com/go/pubsub"
)

func (psClient *Client) Publish(topic string, message Message) error {
	log.Printf("Publishing message to PubSub")

	// Set topic
	t := psClient.PubSubClient.Topic(topic)

	// Get Message data
	data, err := message.GetData()
	if err != nil {
		log.Printf("Could not get data from Message")
		return err
	}

	// Get Message attributes
	attr, err := message.GetAttributes()
	if err != nil {
		log.Printf("Could not get attributes from Message")
		return err
	}

	result := t.Publish(psClient.Ctx, &pubsub.Message{
		Data:       data,
		Attributes: attr,
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(psClient.Ctx)
	if err != nil {
		log.Printf("Get: %v", err)
		return fmt.Errorf("Get: %v", err)
	}

	log.Printf("Published a message; msg ID: %v\n", id)

	psClient.PubSubClient.Close()
	psClient.Ctx.Done()

	return nil
}

package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

func (psClient *Client) Pull(subscription string, handler func(message Message) bool) error {
	log.Printf("Starting subscription %s", subscription)

	// Set Subscription
	sub := psClient.PubSubClient.Subscription(subscription)

	// Setup message flow control
	sub.ReceiveSettings.Synchronous = psClient.Synchronous
	sub.ReceiveSettings.MaxOutstandingMessages = psClient.MaxMessages

	// Receive and handle messages
	err := sub.Receive(psClient.Ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.StartRequestContext(msg.ID)

		log.Printf("Got message: %q\n", string(msg.Data))

		// Create new Message from PubSub Message
		message := NewPubSubMessage(msg)

		// Pass the message to the Handler
		ret := handler(message)

		if ret {
			log.Printf("Successfully processed message: %s\n", msg.ID)
			msg.Ack()
		} else {
			log.Printf("Failed to process message: %s\n", msg.ID)
			msg.Nack()
		}
		log.EndRequestContext(msg.ID)
	})

	if err != nil {
		log.Fatalf("Error processing subscription: %v\n", err)

		return err
	}

	return nil
}

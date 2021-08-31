package pubsub

import "cloud.google.com/go/pubsub"

type Message interface {
	GetData() ([]byte, error)
	GetAttributes() (map[string]string, error)
}

type MessageReceived interface {
	Message
	GetId() (string, error)
}

type PubSubMessage struct {
	Message *pubsub.Message
}

func NewPubSubMessage(msg *pubsub.Message) MessageReceived {
	return &PubSubMessage{Message: msg}
}

func (message PubSubMessage) GetData() ([]byte, error) {
	return message.Message.Data, nil
}

func (message PubSubMessage) GetAttributes() (map[string]string, error) {
	return message.Message.Attributes, nil
}

func (message PubSubMessage) GetId() (string, error) {
	return message.Message.ID, nil
}

package contracts

import (
	"encoding/json"
	"log"
)

type Envelope struct {
	MessageId   string      `json:"id"`
	ClientType  string      `json:"clientType"`
	MessageType string      `json:"messageType"`
	Message     interface{} `json:"message"`
}

type EnvelopeInterface interface {
	GetMessageId() string
	SetMessageId(id string)
	GetClientType() string
	SetClientType(clientType string)
	GetMessageType() string
	SetMessageType(messageType string)
	GetMessage() interface{}
	SetMessage(message interface{})
}

func (envelope Envelope) GetMessageId() string {
	return envelope.MessageId
}

func (envelope *Envelope) SetMessageId(id string) {
	envelope.MessageId = id
}

func (envelope Envelope) GetClientType() string {
	return envelope.ClientType
}

func (envelope *Envelope) SetClientType(clientType string) {
	envelope.ClientType = clientType
}

func (envelope Envelope) GetMessageType() string {
	return envelope.MessageType
}

func (envelope *Envelope) SetMessageType(messageType string) {
	envelope.MessageType = messageType
}

func (envelope Envelope) GetMessage() interface{} {
	return envelope.Message
}

func (envelope *Envelope) SetMessage(message interface{}) {
	envelope.Message = message
}

func (envelope Envelope) GetId() string {
	return envelope.MessageId
}

func (envelope Envelope) GetAttributes() (map[string]string, error) {
	attr := make(map[string]string)

	return attr, nil
}

func (envelope Envelope) GetData() ([]byte, error) {
	data, err := json.Marshal(envelope)

	if err != nil {
		log.Printf("Error parsing envelope into JSON")
		return []byte{}, err
	}

	return data, nil
}

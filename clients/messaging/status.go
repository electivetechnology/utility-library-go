package messaging

import (
	"encoding/json"
	"log"
	"time"
)

const (
	STATUS_CREATED = "created"
)

type Status struct {
	MessageID string    `json:"message_id"` // Message ID uuid (string)
	RequestID string    `json:"request_id"` // Requestor reference
	Status    string    `json:"status"`
	Reason    string    `json:"reason"`
	Requestor string    `json:"requestor"`
	Timestamp time.Time `json:"timestamp"` // Origin timestamp
}

func (status Status) GetId() string {
	return status.MessageID
}

func (status Status) GetAttributes() (map[string]string, error) {
	attr := make(map[string]string)

	return attr, nil
}

func (status Status) GetData() ([]byte, error) {
	data, err := json.Marshal(status)

	if err != nil {
		log.Printf("Error parsing Status into JSON")
		return []byte{}, err
	}

	return data, nil
}

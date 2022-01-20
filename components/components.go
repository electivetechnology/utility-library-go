package components

import (
	"encoding/json"
	"log"
)

type ComponentNotification struct {
	ID       string `json:"id"`
	Callback string `json:"callback"`
}

type ComponentStatusResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (c ComponentNotification) GetData() ([]byte, error) {
	data, err := json.Marshal(c)

	if err != nil {
		log.Printf("Error parsing ComponentNotification into JSON")
		return []byte{}, err
	}

	return data, nil
}

func (c ComponentNotification) GetAttributes() (map[string]string, error) {
	attr := make(map[string]string)

	return attr, nil
}

func (c ComponentNotification) GetId() string {
	return c.ID
}

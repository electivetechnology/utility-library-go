package acl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ExchangePayload struct {
	Organisation string `json:"organisation"`
	Token        string `json:"token"`
}

type ExchangeResponse struct {
	Token string `json:"token"`
}

func (client Client) ExchangeToken(payload *ExchangePayload) (ExchangeResponse, error) {
	// Transform Token struct to json payload
	jsonValue, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, client.BaseUrl+TOKEN_EXCHANGE_URL, bytes.NewBuffer(jsonValue))
	log.Printf("Exchanging token for organisation %s", payload.Organisation)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+payload.Token)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("User-Agent", client.Name)

	// Perform Request
	res, wasRequested, err := client.HandleRequest(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error processing Token exchange: %v", err)
		return ExchangeResponse{}, err
	}

	// Check if the request was actually made
	if !wasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return ExchangeResponse{payload.Token}, nil
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		result := ExchangeResponse{}
		json.Unmarshal(data, &result)

		// Return token
		return result, nil
	}

	return ExchangeResponse{}, errors.New("error exchanging Token for new organisation")
}

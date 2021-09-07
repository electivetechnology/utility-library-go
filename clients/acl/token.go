package acl

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

const (
	TOKEN_TAG_PREFIX = "token_"
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
	request, _ := http.NewRequest(http.MethodPost, client.ApiClient.GetBaseUrl()+TOKEN_EXCHANGE_URL, bytes.NewBuffer(jsonValue))
	log.Printf("Exchanging token for organisation %s", payload.Organisation)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+payload.Token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(payload.Token + payload.Organisation)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error processing Token exchange: %v", err)
		return ExchangeResponse{}, connect.NewInternalError(err.Error())
	}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return ExchangeResponse{payload.Token}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	if res.HttpResponse.StatusCode == http.StatusOK {
		result := ExchangeResponse{}
		json.Unmarshal(data, &result)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return token
		return result, nil
	}

	return ExchangeResponse{}, errors.New("error exchanging Token for new organisation")
}

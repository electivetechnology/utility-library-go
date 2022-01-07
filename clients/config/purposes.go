package config

import (
	"encoding/json"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
)

const (
	PURPOSES_URL       = "/v1/purposes"
	PURPOSE_TAG_PREFIX = "purposes_"
)

type Purpose struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PurposeResponse struct {
	ApiResponse *connect.ApiResponse
	Data        []Purpose
}

func purposeRequest(path string, tagPrefix string, token string, client Client, formatData func(data []byte) []Purpose) (PurposeResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Option", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, purpose evaluation is false
	if err != nil {
		log.Printf("Error getting Option details: %v", err)
		return PurposeResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := PurposeResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return PurposeResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, key)

			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func (client Client) GetPurposes(token string) (PurposeResponse, error) {
	log.Printf("Will request purpose option with purposeId")

	path := client.ApiClient.GetBaseUrl() + PURPOSES_URL

	var formatData = func(data []byte) []Purpose {
		var responseData []Purpose
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return purposeRequest(path, PURPOSE_TAG_PREFIX, token, client, formatData)
}

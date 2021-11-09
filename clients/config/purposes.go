package config

import (
	"encoding/json"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
)

const (
	PUPRPOSES_URL               = "/v1/purposes"
	PUPRPOSE_OPTIONS_TAG_PREFIX = "purposes-options_"
)

type PurposeOptionResponse struct {
	ApiResponse *connect.ApiResponse
	Data        map[string]string
}

func purposeOptionRequest(path string, tagPrefix string, token string, client Client, formatData func(data []byte) map[string]string) (PurposeOptionResponse, error) {
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
		return PurposeOptionResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := PurposeOptionResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return PurposeOptionResponse{}, nil
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
		var responseData map[string]string
		json.Unmarshal(data, &responseData)
		response.Data = responseData
		//response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func (client Client) GetPurposeOption(token string) (PurposeOptionResponse, error) {
	log.Printf("Will request purpose option with purposeId")

	path := client.ApiClient.GetBaseUrl() + PUPRPOSES_URL + "/options"

	var formatData = func(data []byte) map[string]string {
		var responseData map[string]string
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return purposeOptionRequest(path, PUPRPOSE_OPTIONS_TAG_PREFIX, token, client, formatData)
}

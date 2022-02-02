package integrations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/integrations")
}

type IntegrationsResponse struct {
	ApiResponse  *connect.ApiResponse
	Integrations []Integration
}

type Integration struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Organisation string `json:"organisation"`
	Data         []Data `json:"data"`
	OAuthId      string `json:"oauthId"`
}

type Data struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func (client Client) GetIntegrations(token string, filters string) (IntegrationsResponse, error) {
	log.Printf("Will request integrations using filter %s", filters)

	// Generate new path replacer
	r := strings.NewReplacer(":filters", filters)
	path := r.Replace(GET_INTEGRATIONS_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(INTEGRATIONS_TAG_PREFIX + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Integrations details: %v", err)
		return IntegrationsResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := IntegrationsResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return IntegrationsResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		integrations := []Integration{}
		json.Unmarshal(data, &integrations)

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
		response.Integrations = integrations

		return response, nil

	default:
		return response, errors.New(fmt.Sprintf("error getting integrations. Client returned %d", res.HttpResponse.StatusCode))
	}

	return response, nil
}

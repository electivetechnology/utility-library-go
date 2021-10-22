package config

import (
	"encoding/json"
	"errors"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
)

const (
	CHANNELS_URL       = "/v1/channels"
	CHANNEL_TAG_PREFIX = "channels_"
)

type ConfigResponse struct {
	ApiResponse *connect.ApiResponse
	Config      *Config
}

type Config struct {
	Id              string `json:"id"`
	ChannelType     string `json:"channelType"`
	Organisation    string `json:"organisation"`
	Name            string `json:"name"`
	IsConfigEnabled bool   `json:"isConfigEnabled"`
}

func (client Client) GetChannels(channelId string, token string) (ConfigResponse, error) {
	log.Printf("Will request config details for channelId %s and id %d", channelId)

	path := client.ApiClient.GetBaseUrl() + CHANNELS_URL

	request, _ := http.NewRequest(http.MethodGet, path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(CHANNEL_TAG_PREFIX + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Config details: %v", err)
		return ConfigResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := ConfigResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return ConfigResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		config := Config{}
		json.Unmarshal(data, &config)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, CHANNEL_TAG_PREFIX+config.Id)
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Config = &config

		return response, nil

	default:
		return response, nil
	}

	return response, errors.New("error getting config for given vendor")
}

package config

import (
	"encoding/json"
	"net/http"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

const (
	DEFAULTS_CONFIG_URL        = "/v1/default-configs"
	DEFAULT_CONFIGS_TAG_PREFIX = "defaultsConfigs_"
)

type DefaultConfigConfig struct {
	DefaultFooter  string `json:"defaultFooter"`
	DefaultHeader  string `json:"defaultHeader"`
	DefaultWrapper string `json:"defaultWrapper"`
}

type DefaultConfig struct {
	Id           string              `json:"id"`
	Name         string              `json:"name"`
	Config       DefaultConfigConfig `json:"config"`
	Organisation string              `json:"organisation"`
	Visibility   string              `json:"visibility"`
}

type DefaultConfigResponse struct {
	ApiResponse *connect.ApiResponse
	Data        DefaultConfig
}

func defaultConfigRequest(path string, tagPrefix string, idOrName string, token string, client Client, formatData func(data []byte) interface{}) (DefaultConfigResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Config", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Config details: %v", err)
		return DefaultConfigResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := DefaultConfigResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return DefaultConfigResponse{}, nil
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

			if idOrName != "" {
				tags = append(tags, idOrName)
			}

			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		var responseData DefaultConfig
		json.Unmarshal(data, &responseData)
		response.Data = responseData
		//response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func (client Client) GetDefaultConfig(defaultConfigIdOrName string, token string) (DefaultConfigResponse, error) {
	log.Printf("Will request default config with defaultId: %s", defaultConfigIdOrName)

	path := client.ApiClient.GetBaseUrl() + DEFAULTS_CONFIG_URL + "/" + defaultConfigIdOrName

	var formatData = func(data []byte) interface{} {
		var responseData DefaultConfig
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return defaultConfigRequest(path, DEFAULT_CONFIGS_TAG_PREFIX, defaultConfigIdOrName, token, client, formatData)
}

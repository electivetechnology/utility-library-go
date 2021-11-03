package config

import (
	"encoding/json"
	"errors"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
	"time"
)

const (
	CHANNELS_URL             = "/v1/channels"
	CHANNELS_TYPE_URL        = "/v1/channel-types"
	CHANNEL_TAG_PREFIX       = "channels_"
	CHANNEL_TYPES_TAG_PREFIX = "channels-types_"
)

type ChannelResponse struct {
	ApiResponse *connect.ApiResponse
	Data        interface{}
}

type Channel struct {
	Id                  string     `json:"id"`
	Organisation        string     `json:"organisation"`
	Name                string     `json:"name"`
	IsInvitationEnabled string     `json:"isInvitationEnabled"`
	IsAssessmentEnabled string     `json:"isAssessmentEnabled"`
	AssessmentType      string     `json:"assessmentType"`
	ChannelType         string     `json:"channelType"`
	CreatedAt           *time.Time `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
	DeletedAt           *time.Time `json:"deletedAt"`
}

type Channels []Channel

type ChannelType struct {
	Id        string     `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type ChannelTypes []ChannelType

func channelRequest(path string, tagPrefix string, id string, token string, client Client, formatData func(data []byte) interface{}) (ChannelResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Config details: %v", err)
		return ChannelResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := ChannelResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return ChannelResponse{}, nil
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

			if id != "" {
				tags = append(tags, tagPrefix+id)
			}

			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}

	return response, errors.New("error getting config for given vendor")
}

func (client Client) GetChannel(channelId string, token string) (ChannelResponse, error) {
	log.Printf("Will request channel with channelId: %s", channelId)

	path := client.ApiClient.GetBaseUrl() + CHANNELS_URL + "/" + channelId

	var formatData = func(data []byte) interface{} {
		var responseData Channel
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return channelRequest(path, CHANNEL_TAG_PREFIX, channelId, token, client, formatData)
}

func (client Client) GetChannels(token string) (ChannelResponse, error) {
	log.Printf("Will request all channels")

	path := client.ApiClient.GetBaseUrl() + CHANNELS_URL

	var formatData = func(data []byte) interface{} {
		var responseData Channels
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return channelRequest(path, CHANNEL_TAG_PREFIX, "", token, client, formatData)
}

func (client Client) GetChannelType(channelTypeId string, token string) (ChannelResponse, error) {
	log.Printf("Will request channel type with channelId: %s", channelTypeId)

	path := client.ApiClient.GetBaseUrl() + CHANNELS_TYPE_URL + "/" + channelTypeId

	var formatData = func(data []byte) interface{} {
		var responseData ChannelType
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return channelRequest(path, CHANNEL_TYPES_TAG_PREFIX, channelTypeId, token, client, formatData)
}

func (client Client) GetChannelTypes(token string) (ChannelResponse, error) {
	log.Printf("Will request all channel types")

	path := client.ApiClient.GetBaseUrl() + CHANNELS_TYPE_URL

	var formatData = func(data []byte) interface{} {
		var responseData ChannelTypes
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return channelRequest(path, CHANNEL_TYPES_TAG_PREFIX, "", token, client, formatData)
}

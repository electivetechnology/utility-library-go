package config

import (
	"encoding/json"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	ORGANISATION_CONTENTS_URL       = "/v1/organisation-contents"
	ORGANISATION_CONTENT_TAG_PREFIX = "organisation_contents_"
)

type OrganisationContent struct {
	ID           string `json:"id"`
	ContentName  string `json:"content_name"`
	Content      string `json:"content"`
	Organisation string `json:"organisation"`
}

type OrganisationContentResponse struct {
	ApiResponse *connect.ApiResponse
	Data        []OrganisationContent
}

func organisationContentRequest(path string, tagPrefix string, client Client, formatData func(data []byte) []OrganisationContent) (OrganisationContentResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)

	request.Header.Add("Content-Option", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, organisationContent evaluation is false
	if err != nil {
		log.Printf("Error getting Option details: %v", err)
		return OrganisationContentResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := OrganisationContentResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return OrganisationContentResponse{}, nil
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

func (client Client) GetOrganisationContents(filters string, query connect.ApiQuery) (OrganisationContentResponse, error) {
	log.Printf("Will request organisationContents")

	// Generate new path replacer
	r := strings.NewReplacer(":filters", filters)
	path := r.Replace(ORGANISATION_CONTENTS_URL)
	log.Printf("New path generated for request %s", path)

	// Generate new path
	values := url.Values{
		"limit":  []string{strconv.Itoa(query.GetLimit())},
		"offset": []string{strconv.Itoa(query.GetOffset())},
	}

	path = client.ApiClient.GetBaseUrl() + path + "?" + values.Encode()

	var formatData = func(data []byte) []OrganisationContent {
		var responseData []OrganisationContent
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return organisationContentRequest(path, ORGANISATION_CONTENT_TAG_PREFIX, client, formatData)
}

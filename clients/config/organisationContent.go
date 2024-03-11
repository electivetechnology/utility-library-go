package config

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
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

type OrganisationContentsResponse struct {
	ApiResponse *connect.ApiResponse
	Data        []OrganisationContent
}

type OrganisationContentResponse struct {
	ApiResponse *connect.ApiResponse
	Data        OrganisationContent
}

type GenericResponse struct {
	ApiResponse *connect.ApiResponse
	Data        interface{}
}

func setTags(res *connect.ApiResponse, client Client, key string) {
	if !res.WasCached {
		// Save response to cache
		log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

		// Generate tags for cache
		var tags []string
		tags = append(tags, key)

		client.ApiClient.SaveToCache(key, res, tags)
	}
}

func handleRequest(path string, tagPrefix string, client Client) (*connect.ApiResponse, string, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)

	request.Header.Add("Content-Option", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	return res, key, err
}

func organisationContentsRequest(path string, tagPrefix string, client Client, formatData func(data []byte) []OrganisationContent) (OrganisationContentsResponse, error) {
	res, key, err := handleRequest(path, tagPrefix, client)

	// Check for errors, organisationContent evaluation is false
	if err != nil {
		log.Printf("Error getting Option details: %v", err)
		return OrganisationContentsResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := OrganisationContentsResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return OrganisationContentsResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		setTags(res, client, key)

		// Return response
		response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func organisationContentRequest(path string, tagPrefix string, client Client, formatData func(data []byte) OrganisationContent) (OrganisationContentResponse, error) {
	res, key, err := handleRequest(path, tagPrefix, client)

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
		setTags(res, client, key)

		// Return response
		response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func (client Client) GetOrganisationContents(filters string, query connect.ApiQuery) (OrganisationContentsResponse, error) {
	log.Printf("Will request organisationContents")

	// Generate new path
	values := url.Values{
		"limit":  []string{strconv.Itoa(query.GetLimit())},
		"offset": []string{strconv.Itoa(query.GetOffset())},
	}

	path := client.ApiClient.GetBaseUrl() + ORGANISATION_CONTENTS_URL + "?" + values.Encode()

	// Generate new path replacer
	r := strings.NewReplacer(":filters", filters)
	path = r.Replace(path)
	log.Printf("New path generated for request %s", path)

	var formatData = func(data []byte) []OrganisationContent {
		var responseData []OrganisationContent
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return organisationContentsRequest(path, ORGANISATION_CONTENT_TAG_PREFIX, client, formatData)
}

func (client Client) GetByOrganisationContent(organisation string, contentName string) (OrganisationContentResponse, error) {
	log.Printf("Will request organisationContent")

	path := client.ApiClient.GetBaseUrl() + ORGANISATION_CONTENTS_URL +
		"/organisation/" + organisation + "/content/" + contentName

	var formatData = func(data []byte) OrganisationContent {
		var responseData OrganisationContent
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return organisationContentRequest(path, ORGANISATION_CONTENT_TAG_PREFIX, client, formatData)
}

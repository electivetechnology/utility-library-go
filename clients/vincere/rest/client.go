package rest

import (
	"os"
	"strings"
)

type RestClient interface {
	GetApiClient() ApiClient
	SearchApplications(index int, payload *strings.Reader) error
}

type Client struct {
	ApiClient ApiClient
}

func NewRestClient(tenant string) RestClient {
	log.Printf("Createing new vincere rest client")
	// Get Base URL
	url := os.Getenv("VINCERE_REST_BASE_URL")
	if url == "" {
		url = "vincere.io"
	}

	// Get api version
	apiVersion := os.Getenv("VINCERE_REST_API_VERSION")
	if apiVersion == "" {
		apiVersion = "v2"
	}

	// Get apiKey
	apiKey := os.Getenv("VINCERE_API_KEY")

	// Build api base url
	baseUrl := "https://" + tenant + "." + url + "/api/" + apiVersion

	apiClient := BaseClient{BaseUrl: baseUrl, ApiKey: apiKey}

	// Create new Assessments Client
	c := Client{ApiClient: &apiClient}

	return c
}

func (c Client) GetApiClient() ApiClient {
	return c.ApiClient
}

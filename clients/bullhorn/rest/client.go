package rest

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/bullhorn/rest")
}

type RestClient interface {
	GetApiClient() bullhorn.ApiClient
	GetBhRestToken(accessToken string) (*RestToken, error)
	AddRestToken(token *RestToken)
	CreateEntitySubscription(name string, entity string, action string) (*bullhorn.EventsSubscription, error)
	GetCandidate(id int) (*Candidate, error)
}

type Client struct {
	ApiClient bullhorn.ApiClient
}

func NewRestClient() RestClient {
	log.Printf("Createing new bullhorn rest client")
	// Get Base URL
	url := os.Getenv("BULLHORN_REST_BASE_URL")
	if url == "" {
		url = "https://rest.bullhornstaffing.com"
	}

	// Default BHRestToken TTL is seconds
	ttl, _ := strconv.Atoi(os.Getenv("BULLHORN_OAUTH_TTL"))
	if ttl == 0 {
		ttl = 86400 // Default token TTL is set to 24 hours
	}

	apiClient := bullhorn.BaseClient{BaseUrl: url, Ttl: ttl, ApiVersion: "*"}

	// Create new Assessments Client
	c := Client{ApiClient: &apiClient}
	log.Printf("Creates new bullhorn rest client with TTL set to %s", ttl)

	return &c
}

func (client Client) GetApiClient() bullhorn.ApiClient {
	return client.ApiClient
}

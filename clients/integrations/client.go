package integrations

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
)

const (
	CLIENT_NAME             = "Elective:UtilityLibrary:Integrations:0.*"
	INTEGRATIONS_TAG_PREFIX = "integrations_"
	GET_INTEGRATIONS_URL    = "/v1/integrations:filters"
)

type IntegrationsClient interface {
	GetIntegrations(token string, filters string) (IntegrationsResponse, error)
}

type Client struct {
	ApiClient connect.ApiClient
}

func NewClient() IntegrationsClient {
	// Get Base URL
	url := os.Getenv("INTEGRATIONS_HOST")

	if url == "" {
		url = "http://integrations-api"
	}

	// Check if client enabled
	ret := os.Getenv("INTEGRATIONS_CLIENT_ENABLED")
	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse INTEGRATIONS_CLIENT_ENABLED as bool value")
		isEnabled = false
	}

	// Check if redis caching is enabled
	ret = os.Getenv("INTEGRATIONS_CLIENT_REDIS_TTL")
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse INTEGRATIONS_CLIENT_REDIS_TTL as integer value")
		ttl = 0
	}
	log.Printf("Client TTL cache is configured to: %d", ttl)

	apiClient := connect.Client{
		BaseUrl:      url,
		Enabled:      isEnabled,
		Name:         CLIENT_NAME,
		Id:           hash.GenerateHash(12),
		CacheEnabled: false,
	}

	// Setup cache adapter
	apiClient.SetupAdapter(ttl, &apiClient)

	// Create new Candidate Client
	c := Client{ApiClient: &apiClient}

	return &c
}

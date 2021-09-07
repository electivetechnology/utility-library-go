package acl

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	ACL_CLIENT_NAME    = "Elective:UtilityLibrary:ACL:0.*"
	TOKEN_EXCHANGE_URL = "/v1/token/exchange"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/acl")
}

type AclClient interface {
	ExchangeToken(payload *ExchangePayload) (ExchangeResponse, error)
}

type Client struct {
	ApiClient connect.ApiClient
}

func NewClient() *Client {
	// Get Base URL
	url := os.Getenv("ACL_HOST")

	if url == "" {
		url = "http://acl-api"
	}

	// Check if client enabled
	ret := os.Getenv("ACL_CLIENT_ENABLED")
	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse ACL_CLIENT_ENABLED as bool value")
		isEnabled = false
	}

	// Check if redis caching is enabled
	ret = os.Getenv("ACL_CLIENT_REDIS_TTL")
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse ACL_CLIENT_REDIS_TTL as integer value")
		ttl = 0
	}
	log.Printf("Client TTL cache is configured to: %d", ttl)

	apiClient := connect.Client{
		BaseUrl:      url,
		Enabled:      isEnabled,
		Name:         ACL_CLIENT_NAME,
		Id:           hash.GenerateHash(12),
		CacheEnabled: false,
	}

	// Setup cache adapter
	apiClient.SetupAdapter(ttl, &apiClient)

	// Create new Candidate Client
	c := Client{ApiClient: apiClient}

	return &c
}

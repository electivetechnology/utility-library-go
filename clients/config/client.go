package config

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	CLIENT_NAME = "Elective:UtilityLibrary:Config:0.*"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/config")
}

type ConfigClient interface {
	GetChannels(token string) (ChannelResponse, error)
	GetChannel(channelId string, token string) (ChannelResponse, error)
	GetChannelTypes(token string) (ChannelResponse, error)
	GetChannelType(channelTypeId string, token string) (ChannelResponse, error)
	GetDefaultConfig(defaultConfigIdOrName string, token string) (DefaultConfigResponse, error)
	GetPurposes(token string) (PurposeResponse, error)
	GetPurposeTemplateVariables(purposeId string, token string) (PurposeTemplateVariableResponse, error)
	GetByOrganisationContent(organisation string, contentName string) (OrganisationContentResponse, error)
	GetOrganisationContents(filter string, query connect.ApiQuery) (OrganisationContentsResponse, error)
}

type Client struct {
	ApiClient connect.ApiClient
}

func NewClient() ConfigClient {
	// Get Base URL
	url := os.Getenv("CONFIG_HOST")

	if url == "" {
		url = "http://config-api"
	}

	url = "http://config-api-gateway"

	// Check if client enabled
	ret := os.Getenv("CONFIG_CLIENT_ENABLED")
	log.Printf("CONFIG_CLIENT_ENABLED: %d", ret)

	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse CONFIG_CLIENT_ENABLED as bool value")
		isEnabled = false
	}

	isEnabled = true

	// Check if redis caching is enabled
	ret = os.Getenv("CONFIG_CLIENT_REDIS_TTL")
	log.Printf("CONFIG_CLIENT_REDIS_TTL: %d", ret)
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse CONFIG_CLIENT_REDIS_TTL as integer value")
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

	// Create new Config Client
	c := Client{ApiClient: &apiClient}

	return &c
}

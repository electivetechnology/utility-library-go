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
	GetPurposeOption(token string) (PurposeOptionResponse, error)
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

	// Check if client enabled
	ret := os.Getenv("CONFIG_CLIENT_ENABLED")
	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse CONFIG_CLIENT_ENABLED as bool value")
		isEnabled = false
	}

	// Check if redis caching is enabled
	ret = os.Getenv("CONFIG_CLIENT_REDIS_TTL")
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

package candidates

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/cache"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	ACL_CLIENT_NAME              = "Elective:UtilityLibrary:Candidates:0.*"
	GET_CANDIDATE_URL            = "/v1/candidates/:candidate"
	GET_CANDIDATE_FOR_VENDOR_URL = "/v1/candidates/vendor/:vendorName/:vendorId"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/candidates")
}

type CandidatesClient interface {
	GetCandidateByVendor(vendor string, vendorId string, token string) (CandidateResponse, error)
}

type Client struct {
	ApiClient connect.ApiClient
}

func NewClient() CandidatesClient {
	// Get Base URL
	url := os.Getenv("CANDIDATES_HOST")

	if url == "" {
		url = "http://candidates-api"
	}

	// Check if client enabled
	ret := os.Getenv("CANDIDATES_CLIENT_ENABLED")
	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse CANDIDATES_CLIENT_ENABLED as bool value")
		isEnabled = false
	}

	// Check if redis caching is enabled
	ret = os.Getenv("CANDIDATES_CLIENT_REDIS_TTL")
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse CANDIDATES_CLIENT_REDIS_TTL as integer value")
		ttl = 0
	}

	apiClient := connect.Client{
		BaseUrl:      url,
		Enabled:      isEnabled,
		Name:         ACL_CLIENT_NAME,
		Id:           hash.GenerateHash(12),
		CacheEnabled: false,
	}

	if ttl > 0 {
		// Enable caching if ttl is grater than 0
		// Now that redis is enabled we need to set up adapter
		log.Printf("Trying Redis cache adapter")
		adapter := cache.NewCacheAdapter()

		// Set new TTL
		adapter.SetTtl(ttl)

		cache := cache.NewCache(adapter)
		err := cache.Ping()
		if err == nil {
			log.Printf("Client cache has been enabled")
			// Adapter is available, let's configure it
			apiClient.Cache = cache

			// Enable cache
			apiClient.CacheEnabled = true
		} else {
			log.Fatalf("Cache was enabled, however adapter is not available")
		}
	}

	// Create new Candidate Client
	c := Client{ApiClient: apiClient}

	return &c
}

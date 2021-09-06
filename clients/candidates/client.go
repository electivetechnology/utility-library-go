package candidates

import (
	"os"
	"strconv"

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
}

type Client struct {
	BaseUrl        string
	Jwt            string
	IsEnabled      bool
	Name           string
	Id             string
	RedisTTL       int
	IsCacheEnabled bool
}

func NewClient() *Client {
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
	isCacheEnabled := false
	ret = os.Getenv("CANDIDATES_CLIENT_REDIS_TTL")
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse CANDIDATES_CLIENT_REDIS_TTL as integer value")
		ttl = 0
	}

	if ttl > 0 {
		// Enable caching if ttl is grater than 0
		isCacheEnabled = true
	}

	return &Client{
		BaseUrl:        url,
		IsEnabled:      isEnabled,
		Name:           ACL_CLIENT_NAME,
		Id:             hash.GenerateHash(12),
		IsCacheEnabled: isCacheEnabled,
		RedisTTL:       ttl,
	}
}

package assessments

import (
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	HOST_ENV               = "ASSESSMENTS_HOST"
	CLIENT_ENABLED_ENV     = "ASSESSMENTS_CLIENT_ENABLED"
	CACHE_TTL_ENV          = "ASSESSMENTS_CLIENT_REDIS_TTL"
	BASE_URL               = "http://assessments-api"
	CLIENT_NAME            = "Elective:UtilityLibrary:Assessments:0.*"
	GET_JOB_URL            = "/v1/jobs/:job"
	GET_JOB_FOR_VENDOR_URL = "/v1/jobs/vendor/:vendorName/:vendorId"
	GET_INVITATION_URL     = "/v1/invitations/:invitation"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/assessments")
}

type AssessmentsClient interface {
	GetApiClient() connect.ApiClient
	GetJobByVendor(vendor string, vendorId string, token string) (JobResponse, error)
	GetJobById(id string, token string) (JobResponse, error)
	GetInvitationById(id string, token string) (InvitationResponse, error)
}

type Client struct {
	ApiClient connect.ApiClient
}

func NewClient() AssessmentsClient {
	// Get Base URL
	url := os.Getenv(HOST_ENV)

	if url == "" {
		url = BASE_URL
	}

	// Check if client enabled
	ret := os.Getenv(CLIENT_ENABLED_ENV)
	isEnabled, err := strconv.ParseBool(ret)
	if err != nil {
		log.Fatalf("Could not parse %s as bool value", CLIENT_ENABLED_ENV)
		isEnabled = false
	}

	// Check if redis caching is enabled
	ret = os.Getenv(CACHE_TTL_ENV)
	ttl, err := strconv.Atoi(ret)
	if err != nil {
		log.Fatalf("Could not parse %s as integer value", CACHE_TTL_ENV)
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

	// Create new Assessments Client
	c := Client{ApiClient: &apiClient}

	return &c
}

func (client Client) GetApiClient() connect.ApiClient {
	return client.ApiClient
}

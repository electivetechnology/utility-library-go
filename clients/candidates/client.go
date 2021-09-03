package candidates

import (
	"net/http"
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
	log = logger.NewLogger("clients/acl")
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

// HandleRequest takes instance of the http.Request and performs request if client is enabled
// It returns instance of http.Response, bool (true if request was actually made, false if not)
// and error should there be any
func (client Client) HandleRequest(request *http.Request) (*http.Response, bool, error) {
	// Create new Http Client
	c := &http.Client{}

	// Set default headers
	request.Header.Set("User-Agent", client.Name)
	request.Header.Set("X-User-Agent-Id", client.Id)

	if !client.IsEnabled {
		log.Printf("Client is disabled. No request will be made. Returning fake Response")
		return &http.Response{Status: "200 OK", StatusCode: 200}, false, nil
	}
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error handling request: %s %s %v", request.Method, request.URL, err)
		return res, true, err
	}

	return res, true, nil
}

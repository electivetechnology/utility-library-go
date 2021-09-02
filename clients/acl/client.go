package acl

import (
	"net/http"
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/logger"
)

const TOKEN_EXCHANGE_URL = "/v1/token/exchange"

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/acl")
}

type AclClient interface {
	ExchangeToken(payload *ExchangePayload) (ExchangeResponse, error)
}

type Client struct {
	BaseUrl   string
	Jwt       string
	IsEnabled bool
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

	return &Client{BaseUrl: url, IsEnabled: isEnabled}
}

// HandleRequest takes instance of the http.Request and performs request if client is enabled
// It returns instance of http.Response, bool (true if request was actually made, false if not)
// and error should there be any
func (client Client) HandleRequest(request *http.Request) (*http.Response, bool, error) {
	// Create new Http Client
	c := &http.Client{}

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

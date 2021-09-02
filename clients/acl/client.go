package acl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/electivetechnology/utility-library-go/logger"
)

const TOKEN_EXCHANGE_URL = "/v1/token/exchange"

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/oauth")
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

func (client Client) ExchangeToken(payload *ExchangePayload) (ExchangeResponse, error) {
	// Transform Token struct to json payload
	jsonValue, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, client.BaseUrl+TOKEN_EXCHANGE_URL, bytes.NewBuffer(jsonValue))
	log.Printf("Exchanging token for organisation %s", payload.Organisation)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+payload.Token)
	request.Header.Add("Content-Type", "application/json")

	// Perform Request
	res, wasRequested, err := client.HandleRequest(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error processing Token exchange: %v", err)
		return ExchangeResponse{}, err
	}

	// Check if the request was actually made
	if !wasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return ExchangeResponse{payload.Token}, nil
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		result := ExchangeResponse{}
		json.Unmarshal(data, &result)

		// Return token
		return result, nil
	}

	return ExchangeResponse{}, errors.New("error exchanging Token for new organisation")
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

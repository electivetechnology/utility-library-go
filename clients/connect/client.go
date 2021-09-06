package connect

import (
	"log"
	"net/http"
)

type ApiClient interface {
	GetBaseUrl() string
	GetJwt() string
	IsEnabled() bool
	GetName() string
	GetId() string
	GetRedisTTL() int
	IsCacheEnabled() bool
	HandleRequest(request *http.Request) (*ApiResponse, error)
}

type Client struct {
	BaseUrl      string
	Jwt          string
	Enabled      bool
	Name         string
	Id           string
	RedisTTL     int
	CacheEnabled bool
}

// HandleRequest takes instance of the http.Request and performs request if client is enabled
// It returns instance of http.Response, bool (true if request was actually made, false if not)
// and error should there be any
func (client Client) HandleRequest(request *http.Request) (*ApiResponse, error) {
	// Create new Http Client
	c := &http.Client{}

	// Set default headers
	request.Header.Set("User-Agent", client.GetName())
	request.Header.Set("X-User-Agent-Id", client.GetId())

	if !client.IsEnabled() {
		log.Printf("Client is disabled. No request will be made. Returning fake Response")
		return &ApiResponse{
			HttpResponse: &http.Response{Status: "200 OK", StatusCode: 200},
			WasRequested: false,
		}, nil
	}
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error handling request: %s %s %v", request.Method, request.URL, err)
		return &ApiResponse{HttpResponse: res, WasRequested: true}, err
	}

	return &ApiResponse{HttpResponse: res, WasRequested: true}, nil
}

func (client Client) GetBaseUrl() string {
	return client.BaseUrl
}

func (client Client) IsEnabled() bool {
	return client.Enabled
}

func (client Client) GetJwt() string {
	return client.Jwt
}

func (client Client) GetName() string {
	return client.Name
}

func (client Client) GetId() string {
	return client.Id
}

func (client Client) GetRedisTTL() int {
	return client.RedisTTL
}

func (client Client) IsCacheEnabled() bool {
	return client.CacheEnabled
}

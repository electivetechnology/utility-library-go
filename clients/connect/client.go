package connect

import (
	"io/ioutil"
	"net/http"

	"github.com/electivetechnology/utility-library-go/cache"
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/connect")
}

type ApiClient interface {
	GetBaseUrl() string
	GetJwt() string
	IsEnabled() bool
	GetName() string
	GetId() string
	GetRedisTTL() int
	IsCacheEnabled() bool
	HandleRequest(request *http.Request, key string) (*ApiResponse, error)
	SaveToCache(key string, response *ApiResponse, tags []string) error
	GenerateKey(key string) string
}

type Client struct {
	BaseUrl      string
	Jwt          string
	Enabled      bool
	Name         string
	Id           string
	RedisTTL     int
	CacheEnabled bool
	Cache        *cache.Cache
}

// HandleRequest takes instance of the http.Request and performs request if client is enabled
func (client Client) HandleRequest(request *http.Request, key string) (*ApiResponse, error) {
	log.Printf("Handling request %s: %s", request.Method, request.RequestURI)
	// Create new Http Client
	c := &http.Client{}

	// Set default headers
	request.Header.Set("User-Agent", client.GetName())
	request.Header.Set("X-User-Agent-Id", client.GetId())

	if !client.IsEnabled() {
		log.Printf("Client is disabled. No request will be made. Returning fake Response")
		return &ApiResponse{
			HttpResponse: &HttpResponse{Status: "200 OK", StatusCode: 200},
			WasRequested: false,
		}, nil
	}

	// Get cached
	log.Printf("Checking cached response is available for this request")

	res, err := client.GetCached(request, key)
	if err != nil {
		// Make request as no result available in cache
		ret, err := c.Do(request)

		// Check for errors, default evaluation is false
		if err != nil {
			log.Printf("Error handling request: %s %s %v", request.Method, request.URL, err)
			return &ApiResponse{HttpResponse: &HttpResponse{Status: ret.Status, StatusCode: ret.StatusCode}, WasRequested: true}, err
		}

		// read all response body
		data, _ := ioutil.ReadAll(ret.Body)

		res = &ApiResponse{HttpResponse: &HttpResponse{Status: ret.Status, StatusCode: ret.StatusCode, Body: data}, WasRequested: true}

		// mark response as not cached
		res.WasCached = false
	} else {
		// indicate that response was from cache
		res.WasCached = true
	}

	return res, nil
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

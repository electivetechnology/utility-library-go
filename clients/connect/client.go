package connect

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	Cache        *cache.Cache
}

// HandleRequest takes instance of the http.Request and performs request if client is enabled
// It returns instance of http.Response, bool (true if request was actually made, false if not)
// and error should there be any
func (client Client) HandleRequest(request *http.Request) (*ApiResponse, error) {
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
	res, err := client.GetCached(request)
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

		if client.IsCacheEnabled() {
			// Save to cache
			client.SaveToCache(request, res)
		}
	}

	return res, nil
}

func (client Client) SaveToCache(request *http.Request, response *ApiResponse) error {
	if client.IsCacheEnabled() {
		// Save  result to cache
		log.Printf("Saving to cache result for request %s: %s", request.Method, request.RequestURI)

		key := client.GetCacheKey(request)
		log.Printf("Cache key for this request is: %s", key)

		data, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Could not encode Response into JSON %v", err)
			return err
		}

		client.Cache.Set(key, data)
	}

	return nil
}

func (client Client) GetCached(request *http.Request) (*ApiResponse, error) {
	if client.IsCacheEnabled() {
		// Cache is enabled check if redis is available
		log.Printf("Getting cached result for request %s: %s", request.Method, request.RequestURI)

		key := client.GetCacheKey(request)
		log.Printf("Cache key for this request is: %s", key)

		data := client.Cache.Get(key)
		var res ApiResponse

		if err := json.Unmarshal(data, &res); err != nil {
			log.Printf("Failed to decode cached data %v", err)

			return &ApiResponse{}, errors.New("failed to decode cached data")
		}

		return &res, nil
	}

	log.Printf("Cache is disabled for request %s: %s", request.Method, request.RequestURI)

	return &ApiResponse{}, errors.New("response not available in cache")
}

func (client Client) GetCacheKey(request *http.Request) string {
	h := md5.New()
	h.Write([]byte(request.Host))
	h.Write([]byte(request.Method))
	h.Write([]byte(request.RequestURI))

	hash := h.Sum(nil)

	return hex.EncodeToString(hash[:])
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

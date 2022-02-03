package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
	"github.com/electivetechnology/utility-library-go/hash"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	HOST_ENV                = "OAUTH_HOST"
	CLIENT_ENABLED_ENV      = "OAUTH_CLIENT_ENABLED"
	CACHE_TTL_ENV           = "OAUTH_CLIENT_REDIS_TTL"
	BASE_URL                = "http://oauth-api"
	CLIENT_NAME             = "Elective:UtilityLibrary:Oauth:0.*"
	AUTH_TOKEN_URL          = "/v1/oauth2/authorizations/:state/token"
	ACCESS_TOKEN_TAG_PREFIX = "access_token"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/oauth")
}

type OAuthClient interface {
	GetToken(auth Authorization) (Token, error)
	RefreshToken(Token Token, clientId string, clientSecret string) (Token, error)
	Refresh(refreshToken string) (Token, error)
	GetAccessToken(state string, jwt string) (TokenResponse, error)
}

type Client struct {
	BaseUrl string
	Jwt     string
}

type ApiClient struct {
	ApiClient connect.ApiClient
}

func NewOAuthClient() OAuthClient {
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
	c := ApiClient{ApiClient: &apiClient}

	return &c
}

func NewClient(jwt string) *Client {
	// Get Base URL
	url := os.Getenv("OAUTH_HOST")

	if url == "" {
		url = "http://oauth"
	}

	return &Client{BaseUrl: url, Jwt: jwt}
}

func (client *ApiClient) GetToken(auth Authorization) (Token, error) {
	return &AccessToken{}, nil
}

func (client *ApiClient) RefreshToken(Token Token, clientId string, clientSecret string) (Token, error) {
	return &AccessToken{}, nil
}

func (client *ApiClient) Refresh(refreshToken string) (Token, error) {
	return &AccessToken{}, nil
}

func (client *ApiClient) GetAccessToken(state string, jwt string) (TokenResponse, error) {
	log.Printf("Sending request to Oauth to get new Token for Auth %s", state)

	// Generate new path replacer
	r := strings.NewReplacer(":state", state)
	path := r.Replace(AUTH_TOKEN_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+jwt)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(ACCESS_TOKEN_TAG_PREFIX + path)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting AccessToken details: %v", err)
		return Response{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	log.Printf("Got response from server: %v", res)
	response := Response{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return Response{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		accessToken := AccessToken{}
		json.Unmarshal(data, &accessToken)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, ACCESS_TOKEN_TAG_PREFIX+accessToken.GetAccessToken())
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Token = &accessToken

		return response, nil

	case http.StatusNotFound:
		msg := fmt.Sprintf("Could not find AccessToken: %s", state)
		return response, errors.New(msg)
	default:
		msg := fmt.Sprintf("Could not get AccessToken details")
		return response, errors.New(msg)
	}
}

func (client *Client) GetToken(auth string) (Token, error) {
	return client.GetAccessToken(auth)
}

func (client *Client) GetAccessToken(auth string) (*AccessToken, error) {
	log.Printf("Sending request to Oauth to get new Token for Auth %s", auth)

	// Prepare Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, client.BaseUrl+"/v1/oauth2/authorizations/"+auth+"/token", nil)
	r.Header.Add("Authorization", "Bearer "+client.Jwt)

	// Send Request
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Access token: %v", err)

		return &AccessToken{}, errors.New("error getting Access token")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		token := &AccessToken{}
		json.Unmarshal(data, &token)

		// Return token
		return token, nil
	}

	// If we got here there was some kind of error with exchange

	return &AccessToken{}, nil
}

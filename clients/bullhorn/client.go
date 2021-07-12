package bullhorn

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/electivetechnology/utility-library-go/clients/oauth"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	AUTH_AUTHORIZATION_URL = "/oauth/authorize"
	AUTH_TOKEN_URL         = "/oauth/token"
	REST_LOGIN_URL         = "/rest-services/login"
	REST_PUT_SUBSCRIPTION  = "event/subscription/"
)

var log logger.Logging

type OAuthClient struct {
	BaseUrl string
}

type RestClient struct {
	BaseUrl    string
	Ttl        int
	ApiVersion string
	RestToken  string
}

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/bullhorn")
}

func NewOAuthClient() *OAuthClient {
	// Get Base URL
	url := os.Getenv("BULLHORN_OAUTH_BASE_URL")

	if url == "" {
		url = "https://auth.bullhornstaffing.com"
	}

	return &OAuthClient{BaseUrl: url}
}

func NewRestClient() *RestClient {
	// Get Base URL
	url := os.Getenv("BULLHORN_REST_BASE_URL")
	if url == "" {
		url = "https://rest.bullhornstaffing.com"
	}

	// Default BHRestToken TTL is seconds
	ttl, _ := strconv.Atoi(os.Getenv("BULLHORN_OAUTH_TTL"))
	if ttl == 0 {
		ttl = 86400 // Default token TTL is set to 24 hours
	}

	return &RestClient{BaseUrl: url, Ttl: ttl, ApiVersion: "*"}
}

func (client *RestClient) AddRestToken(token *RestToken) *RestClient {
	client.BaseUrl = token.RestUrl
	client.RestToken = token.BhRestToken

	return client
}

func (client *OAuthClient) GetToken(auth oauth.Authorization) (oauth.Token, error) {
	// Get Redirect URLS
	redirectUrl := os.Getenv("OAUTH_REDIRECT_BASE_URL")

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"grant_type":    []string{oauth.GRANT_TYPE_AUTH_CODE},
		"code":          []string{auth.GetCode()},
		"client_id":     []string{auth.GetClientId()},
		"client_secret": []string{auth.GetClientSecret()},
		"redirect_uri":  []string{redirectUrl},
	}

	log.Printf("Sending following data to bullhorn for exchange: %v", values)

	// Perform Request
	res, err := http.PostForm(client.BaseUrl+AUTH_TOKEN_URL, values)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error exchanging Authorization for Access token: %v\n", err)
		return &Token{}, errors.New("error exchanging Authorization for Access token")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		token := Token{}
		json.Unmarshal(data, &token)

		// Set ExpiresAt
		t := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
		token.ExpiresAt = &t

		// Return token
		return token, nil
	}

	// If we got here there was some kind of error with exchange

	return Token{}, errors.New("error exchanging Authorization for Access token")
}

func (client *OAuthClient) Refresh(refreshToken string) (oauth.Token, error) {
	// If we got here there was some kind of error with exchange

	return Token{}, errors.New("error exchanging Refresh Token for Access token")
}

func (client *OAuthClient) RefreshToken(token oauth.Token, clientId string, clientSecret string) (oauth.Token, error) {
	// If we got here there was some kind of error with exchange
	// Transform Token struct to RefreshToken payload
	// Set URL parameters on declaration
	values := url.Values{
		"grant_type":    []string{oauth.GRANT_TYPE_REFRESH_TOKEN},
		"refresh_token": []string{token.GetRefreshToken()},
		"client_id":     []string{clientId},
		"client_secret": []string{clientSecret},
	}

	log.Printf("Sending following data to bullhorn for refresh: %v", values)

	// Perform Request
	res, err := http.PostForm(client.BaseUrl+AUTH_TOKEN_URL, values)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error exchanging Authorization for Access token: %v\n", err)
		return &Token{}, errors.New("error exchanging Authorization for Access token")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		token := Token{}
		json.Unmarshal(data, &token)

		// Set ExpiresAt
		t := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
		token.ExpiresAt = &t

		// Return token
		return token, nil
	}

	// If we got here there was some kind of error with exchange

	return Token{}, errors.New("error exchanging Token for Access token")
}

func (client *RestClient) GetBhRestToken(accessToken string) (*RestToken, error) {
	return client.Login(accessToken)
}

func (client *RestClient) Login(accessToken string) (*RestToken, error) {
	log.Printf("Will login to Bullhorn and get BhRestToken")

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"version":      []string{client.ApiVersion},
		"access_token": []string{accessToken},
		"ttl":          []string{strconv.Itoa(client.Ttl / 60)}, // We need to covert seconds to minutes
	}

	log.Printf("Sending following data to bullhorn for login: %v", values)

	// Perform Request
	res, err := http.PostForm(client.BaseUrl+REST_LOGIN_URL, values)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &RestToken{}, errors.New("error exchanging Authorization for Access token")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		token := RestToken{}
		json.Unmarshal(data, &token)

		// Set ExpiresAt
		t := time.Now().Add(time.Duration(client.Ttl))
		token.ExpiresAt = &t

		token.Ttl = client.Ttl

		// Return token
		return &token, nil
	}

	// If we got here there was some kind of error with exchange

	return &RestToken{}, errors.New("error exchanging Access token for BH Rest Token")
}

func (client *RestClient) CreateEntitySubscription(name string, entity string, action string) (*EventsSubscription, error) {
	var entities []string
	var actions []string

	// Add entity
	entities = append(entities, entity)

	// Add actions
	actions = append(actions, action)

	// Forward request to multi entity/action handler
	return client.CreateSubscription(name, "entity", entities, actions)
}

func (client *RestClient) CreateSubscription(name string, subType string, entities []string, actions []string) (*EventsSubscription, error) {
	log.Printf("Will CreateSubscription subscription for the following entities %s", strings.Join(entities, ","))

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"type":        []string{"entity"},
		"names":       entities,
		"eventTypes":  actions,
		"BhRestToken": []string{client.RestToken},
	}

	log.Printf("Sending following data to bullhorn for new subscription: %v", values)

	// generate URL for request
	requestUrl, err := GenerateURL(client.BaseUrl, REST_PUT_SUBSCRIPTION+name, values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &EventsSubscription{}, errors.New("error creating new subscription")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPut, requestUrl, nil) // URL-encoded payload

	log.Printf("Adding Header for bullhorn authorization %s", client.RestToken)

	// Print request details
	log.Printf("Request details: %v", r)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &EventsSubscription{}, errors.New("error creating new subscription")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		subscription := EventsSubscription{}
		json.Unmarshal(data, &subscription)

		// Return token
		return &subscription, nil
	}

	// If we got here there was some kind of error with exchange

	return &EventsSubscription{}, errors.New("error sending request to bullhorn")
}

package bullhorn

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/electivetechnology/utility-library-go/clients/oauth"
	"github.com/electivetechnology/utility-library-go/logger"
)

const (
	AUTH_AUTHORIZATION_URL = "/oauth/authorize"
	AUTH_TOKEN_URL         = "/oauth/token"
)

var log logger.Logging

type OAuthClient struct {
	BaseUrl string
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

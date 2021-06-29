package oauth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/electivetechnology/utility-library-go/logger"
)

const AUTH_TOKEN_URL = "/v1/oauth2/authorizations/:state/token"

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("clients/oauth")
}

type OAuthClient interface {
	GetToken(auth Authorization) (Token, error)
	RefreshToken(Token Token, clientId string, clientSecret string) (Token, error)
	Refresh(refreshToken string) (Token, error)
}

type Client struct {
	BaseUrl string
	Jwt     string
}

func NewClient(jwt string) *Client {
	// Get Base URL
	url := os.Getenv("OAUTH_HOST")

	if url == "" {
		url = "http://oauth"
	}

	return &Client{BaseUrl: url, Jwt: jwt}
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

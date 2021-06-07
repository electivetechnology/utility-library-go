package google

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/electivetechnology/utility-library-go/clients/oauth"
)

const (
	AUTH_TOKEN_URL = "/token"
)

type OAuthClient struct {
	BaseUrl string
	Host    string
}

func NewOAuthClient() *OAuthClient {
	// Get Base URL
	url := os.Getenv("GOOGLE_OAUTH2_BASE_URL")

	if url == "" {
		url = "https://oauth2.googleapis.com"
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

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, client.BaseUrl+AUTH_TOKEN_URL, strings.NewReader(values.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	//res, err := http.PostForm(client.BaseUrl+AUTH_TOKEN_URL, values)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error exchanging Authorization for Access token: %v\n", err)
		return &Token{}, errors.New("error exchanging Authorization for Access token")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

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

	return &Token{}, errors.New("error exchanging Authorization for Access token")
}

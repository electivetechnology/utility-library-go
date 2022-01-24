package vincere

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
)

const (
	AUTH_AUTHORIZATION_URL = "/oauth2/authorize"
	AUTH_TOKEN_URL         = "/oauth2/token"
)

type OAuthClient struct {
	BaseUrl string
}

func NewOAuthClient() *OAuthClient {
	// Get Base URL
	url := os.Getenv("VINCERE_OAUTH2_BASE_URL")

	if url == "" {
		url = "https://id.vincere.io"
	}

	return &OAuthClient{BaseUrl: url}
}

func (client *OAuthClient) GetToken(auth oauth.Authorization) (oauth.Token, error) {
	// Get Redirect URLS
	redirectUrl := os.Getenv("VINCERE_OAUTH2_BASE_URL")

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"grant_type":    []string{oauth.GRANT_TYPE_AUTH_CODE},
		"code":          []string{auth.GetCode()},
		"client_id":     []string{auth.GetClientId()},
		"client_secret": []string{auth.GetClientSecret()},
		"redirect_uri":  []string{redirectUrl},
	}

	log.Printf("Sending following data to google for exchange: %v", values)

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, client.BaseUrl+AUTH_TOKEN_URL, strings.NewReader(values.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

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

	return &Token{}, errors.New("error exchanging Authorization for Access token")
}

// Refresh is not implemented by Vincere
func (client *OAuthClient) Refresh(refreshToken string) (oauth.Token, error) {
	// If we got here there was some kind of error with exchange

	return Token{}, errors.New("error exchanging Refresh Token for Access token")
}

func (client *OAuthClient) RefreshToken(token oauth.Token, clientId string, clientSecret string) (oauth.Token, error) {
	// Set URL parameters on declaration
	values := url.Values{
		"grant_type":    []string{oauth.GRANT_TYPE_REFRESH_TOKEN},
		"refresh_token": []string{token.GetRefreshToken()},
		"client_id":     []string{clientId},
	}

	log.Printf("Sending following data to vincere for refresh: %v", values)

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, client.BaseUrl+AUTH_TOKEN_URL, strings.NewReader(values.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error refreshing token: %v\n", err)
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
		tkn := Token{}
		json.Unmarshal(data, &tkn)

		// Set ExpiresAt
		t := time.Now().Add(time.Second * time.Duration(tkn.ExpiresIn))
		tkn.ExpiresAt = &t

		// Add existing refresh token (vincare uses one refresh token)
		tkn.RefreshToken = token.GetRefreshToken()

		// Return token
		return tkn, nil
	}

	// If we got here there was some kind of error with exchange

	return Token{}, errors.New("error exchanging Token for Access token")
}

package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	REST_LOGIN_URL = "/rest-services/login"
)

func (client *Client) Login(accessToken string) (*RestToken, error) {
	log.Printf("Will login to Bullhorn and get BhRestToken")

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"version":      []string{client.GetApiClient().GetApiVersion()},
		"access_token": []string{accessToken},
		"ttl":          []string{strconv.Itoa(client.ApiClient.GetTtl() / 60)}, // We need to covert seconds to minutes
	}

	log.Printf("Sending following data to bullhorn for login: %v", values)

	// Perform Request
	res, err := http.PostForm(client.ApiClient.GetBaseUrl()+REST_LOGIN_URL, values)

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
		t := time.Now().Add(time.Duration(client.ApiClient.GetTtl()))
		token.ExpiresAt = &t

		token.Ttl = client.ApiClient.GetTtl()

		// Return token
		return &token, nil
	}

	// If we got here there was some kind of error with exchange

	return &RestToken{}, errors.New("error exchanging Access token for BH Rest Token")
}

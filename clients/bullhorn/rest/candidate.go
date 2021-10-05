package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
)

const (
	REST_GET_CANDIDATE = "entity/Candidate/"
)

type CandidateResponse struct {
	Data Candidate `json:"data"`
}

type Candidate struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Status    string  `json:"status"`
	Address   Address `json:"address"`
	Mobile    string  `json:"mobile"`
}

type Address struct {
	City        string `json:"city"`
	CountryName string `json:"countryName"`
}

func (client *Client) GetCandidate(id int) (*Candidate, error) {
	log.Printf("Will Get Candidate details with bullhorn ID %d", id)

	// Set URL parameters on declaration
	values := url.Values{
		"fields":      []string{"firstName,lastName,address,email,phone,status,mobile"},
		"BhRestToken": []string{client.GetApiClient().GetRestToken()},
	}

	log.Printf("Sending following data to bullhorn to get Candidate: %v", values)

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), REST_GET_CANDIDATE+strconv.Itoa(id), values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &Candidate{}, errors.New("error creating getting Candidate details")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, requestUrl, nil) // URL-encoded payload

	// Print request details
	log.Printf("Request details: %v", r)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &Candidate{}, errors.New("error creating new subscription")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		ret := CandidateResponse{}
		json.Unmarshal(data, &ret)

		// Return token
		return &ret.Data, nil
	}

	return &Candidate{}, nil
}

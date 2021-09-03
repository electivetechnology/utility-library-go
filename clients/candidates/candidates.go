package candidates

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type CandidateResponse struct {
	Id                string `json:"id"`
	Email             string `json:"email"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Phone             string `json:"phone"`
	Organisation      string `json:"organisation"`
	PrimaryLanguage   string `json:"primaryLanguage"`
	SecondaryLanguage string `json:"secondaryLanguage"`
	TertiaryLanguage  string `json:"tertiaryLanguage"`
	Title             string `json:"title"`
	Gender            string `json:"gender"`
}

func (client Client) GetCandidateByVendor(vendor string, vendorId string, token string) (CandidateResponse, error) {
	log.Printf("Will request candidate details for vendor %s and id %s", vendor, vendorId)

	// Generate new path replacer
	r := strings.NewReplacer(":vendor", vendor, ":vendorId", vendorId)
	path := r.Replace(GET_CANDIDATE_FOR_VENDOR_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.BaseUrl+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Perform Request
	res, wasRequested, err := client.HandleRequest(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Candidate details: %v", err)
		return CandidateResponse{}, err
	}

	// Check if the request was actually made
	if !wasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return CandidateResponse{}, nil
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		result := CandidateResponse{}
		json.Unmarshal(data, &result)

		// Return token
		return result, nil
	}

	return CandidateResponse{}, errors.New("error getting candidate for given vendor")
}

package candidates

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

const (
	CANDIDATE_TAG_PREFIX = "candidates_"
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
	r := strings.NewReplacer(":vendorName", vendor, ":vendorId", vendorId)
	path := r.Replace(GET_CANDIDATE_FOR_VENDOR_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Generate tags for cache
	var tags []string

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, tags)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Candidate details: %v", err)
		return CandidateResponse{}, connect.NewInternalError(err.Error())
	}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return CandidateResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		result := CandidateResponse{}
		json.Unmarshal(data, &result)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache")

			// Add Id to tags
			tags = append(tags, CANDIDATE_TAG_PREFIX+result.Id)
			client.ApiClient.SaveToCache(request, res, tags)
		}

		// Return token
		return result, nil

	case http.StatusNotFound:
		// Return 404
		e := connect.NewApiError("candidate not found")
		e.Status = 404
		return CandidateResponse{}, e
	}

	return CandidateResponse{}, errors.New("error getting candidate for given vendor")
}

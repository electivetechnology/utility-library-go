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
	ApiResponse *connect.ApiResponse
	Candidate   *Candidate
}

type Candidate struct {
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
	log.Printf("Will request candidate details for vendor %s and id %d", vendor, vendorId)

	// Generate new path replacer
	r := strings.NewReplacer(":vendorName", vendor, ":vendorId", vendorId)
	path := r.Replace(GET_CANDIDATE_FOR_VENDOR_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(CANDIDATE_TAG_PREFIX + path)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Candidate details: %v", err)
		return CandidateResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := CandidateResponse{ApiResponse: res}

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

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		candidate := Candidate{}
		json.Unmarshal(data, &candidate)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, CANDIDATE_TAG_PREFIX+candidate.Id)
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Candidate = &candidate

		return response, nil

	default:
		return response, nil
	}

	return response, errors.New("error getting candidate for given vendor")
}

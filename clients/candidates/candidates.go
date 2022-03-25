package candidates

import (
	"bytes"
	"encoding/hex"
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

type CandidateVendorResponse struct {
	ApiResponse     *connect.ApiResponse
	CandidateVendor *CandidateVendor
}

type CandidateVendor struct {
	Id          string `json:"id"`
	VendorId    string `json:"vendorId"`
	Vendor      string `json:"vendor"`
	CandidateId string `json:"candidateId"`
}

type VendorPayload struct {
	Vendor   string `json:"vendor"`
	VendorId string `json:"vendorId"`
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
	CvText            string `json:"cvText"`
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
	key := client.ApiClient.GenerateKey(CANDIDATE_TAG_PREFIX + path + token)

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

func (client Client) GetCandidateById(id string, token string) (CandidateResponse, error) {
	log.Printf("Will request candidate details for id %d", id)

	// Generate new path replacer
	r := strings.NewReplacer(":candidate", id)
	path := r.Replace(GET_CANDIDATE_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(CANDIDATE_TAG_PREFIX + path + token)

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

func (client Client) PutCandidate(payload []byte, token string) (CandidateResponse, error) {
	log.Printf("Sending request to Create Candidate")

	// convert byte slice to io.Reader
	reader := bytes.NewReader(payload)

	request, _ := http.NewRequest(http.MethodPut, client.ApiClient.GetBaseUrl()+PUT_CANDIDATE_URL, reader)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	pld := hex.EncodeToString(payload)

	// Get key
	key := client.ApiClient.GenerateKey(CANDIDATE_TAG_PREFIX + PUT_CANDIDATE_URL + pld + token)

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

	case http.StatusCreated:
		candidate := Candidate{}
		json.Unmarshal(data, &candidate)

		// Return response
		response.Candidate = &candidate

		return response, nil
	default:
		return response, nil
	}
}

func (client Client) AddCandidateVendor(vendor string, vendorId string, token string) (CandidateVendorResponse, error) {
	log.Printf("Sending request to Create Candidate")

	// vendor and id to json payload
	jsonValue, _ := json.Marshal(VendorPayload{Vendor: vendor, VendorId: vendorId})

	request, _ := http.NewRequest(http.MethodPost, client.ApiClient.GetBaseUrl()+ADD_CANDIDATE_VENDOR, bytes.NewBuffer(jsonValue))

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(CANDIDATE_TAG_PREFIX + ADD_CANDIDATE_VENDOR + vendor + vendorId + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Candidate details: %v", err)
		return CandidateVendorResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := CandidateVendorResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return CandidateVendorResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		vendor := CandidateVendor{}
		json.Unmarshal(data, &vendor)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, CANDIDATE_TAG_PREFIX+vendor.Vendor+vendor.Id)
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.CandidateVendor = &vendor

		return response, nil

	case http.StatusCreated:
		vendor := CandidateVendor{}
		json.Unmarshal(data, &vendor)

		// Return response
		response.CandidateVendor = &vendor

		return response, nil
	case http.StatusConflict:
		log.Printf("Vendor already exists")
		return response, nil
	default:
		return response, nil
	}
}

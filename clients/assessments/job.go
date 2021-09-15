package assessments

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

const (
	JOB_TAG_PREFIX = "jobs_"
)

type JobResponse struct {
	ApiResponse *connect.ApiResponse
	Job         *Job
}

type Job struct {
	Id               string   `json:"id"`
	Title            string   `json:"title"`
	CandidateSummary string   `json:"candidateSummary"`
	Brief            string   `json:"brief"`
	Type             string   `json:"type"`
	Currency         string   `json:"currency"`
	Salary           string   `json:"salary"`
	SalaryUnit       string   `json:"salaryUnit"`
	Location         string   `json:"location"`
	Keywords         []string `json:"keywords"`
	ClientId         string   `json:"clientId"`
	ClientName       string   `json:"clientName"`
	Status           string   `json:"status"`
	Headline         string   `json:"headline"`
	Notes            string   `json:"notes"`
}

func (client Client) GetJobByVendor(vendor string, vendorId string, token string) (JobResponse, error) {
	log.Printf("Will request job details for vendor %s and id %s", vendor, vendorId)

	// Generate new path replacer
	r := strings.NewReplacer(":vendorName", vendor, ":vendorId", vendorId)
	path := r.Replace(GET_JOB_FOR_VENDOR_URL)
	log.Printf("New path generated for request %s", path)

	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey("")

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error getting Job details: %v", err)
		return JobResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	log.Printf("Got response from server: %v", res)
	response := JobResponse{ApiResponse: res}

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return JobResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		job := Job{}
		json.Unmarshal(data, &job)

		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, JOB_TAG_PREFIX+job.Id)
			tags = append(tags, key)
			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Job = &job

		return response, nil

	default:
		return response, nil
	}

	return JobResponse{}, nil
}

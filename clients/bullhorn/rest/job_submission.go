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
	REST_GET_JOB_SUBMISSION = "entity/JobSubmission/"
)

type JobSubmissionResponse struct {
	Data JobSubmission `json:"data"`
}

type JobSubmission struct {
	Status    string    `json:"status"`
	Candidate Candidate `json:"candidate"`
	JobOrder  JobOrder  `json:"jobOrder"`
}

type JobOrder struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (client *Client) GetJobSubmission(id int) (*JobSubmission, error) {
	log.Printf("Will Get JobSubmission details with ID %d", id)

	// Set URL parameters on declaration
	values := url.Values{
		"fields":      []string{"status,candidate,jobOrder"},
		"BhRestToken": []string{client.GetApiClient().GetRestToken()},
	}

	log.Printf("Sending following data to bullhorn to get JobSubmission: %v", values)

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), REST_GET_JOB_SUBMISSION+strconv.Itoa(id), values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &JobSubmission{}, errors.New("error getting JobSubmission details")
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
		return &JobSubmission{}, errors.New("error creating new subscription")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		ret := JobSubmissionResponse{}
		json.Unmarshal(data, &ret)

		// Return token
		return &ret.Data, nil
	}

	return &JobSubmission{}, nil
}

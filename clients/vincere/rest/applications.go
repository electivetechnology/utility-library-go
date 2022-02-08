package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
)

const (
	APPLICATIONS_SEARCH_URL = "/application/filter"
)

type Applications struct {
	Content []Application
	Index   int  `json:"slice_index"`
	Count   int  `json:"num_of_elements"`
	Last    bool `json:"last"`
}

type Application struct {
	Id          int `json:"id"`
	CandidateId int `json:"candidate_id"`
	JobId       int `json:"job_id"`
}

func (client Client) SearchApplications(index int, payload []byte, token string) (Applications, error) {
	applications := Applications{}
	// Set URL parameters on declaration
	values := url.Values{
		"index": []string{strconv.Itoa(index)},
	}

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), APPLICATIONS_SEARCH_URL, values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return applications, errors.New("could not generate URL for application search")
	}

	// convert byte slice to io.Reader
	reader := bytes.NewReader(payload)

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, requestUrl, reader) // URL-encoded payload

	// Add headers
	r.Header.Add("id-token", token)
	r.Header.Add("x-api-key", client.GetApiClient().GetApiKey())
	r.Header.Add("Content-Type", "application/json")

	log.Printf("Sending request with payload %v and x-api-key: %s and id-token %s", payload, client.GetApiClient().GetApiKey(), token)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error Searching Applications: %v\n", err)
		return applications, errors.New("error searching applications")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	switch res.StatusCode {
	case http.StatusOK:
		json.Unmarshal(data, &applications)

		return applications, nil
	default:
		msg := fmt.Sprintf("Could not get Applications details")
		return applications, errors.New(msg)
	}
}

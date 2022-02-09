package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	CANDIDATE_TAG_PREFIX = "vincere_candidate"
	CANDIDATE_GET_URL    = "/candidate/:id"
)

type Candidate struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (client Client) GetCandidate(id int, token string) (Candidate, error) {
	candidate := Candidate{}
	// Generate new path replacer
	r := strings.NewReplacer(":id", strconv.Itoa(id))
	path := r.Replace(CANDIDATE_GET_URL)
	log.Printf("New path generated for request %s", path)

	c := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, client.ApiClient.GetBaseUrl()+path, nil)

	// Add headers
	request.Header.Add("id-token", token)
	request.Header.Add("x-api-key", client.GetApiClient().GetApiKey())
	request.Header.Add("Content-Type", "application/json")

	// Perform Request
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error Searching Applications: %v\n", err)
		return candidate, errors.New("error searching applications")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		json.Unmarshal(data, &candidate)

		return candidate, nil
	default:
		msg := fmt.Sprintf("Could not get Applications details")
		return candidate, errors.New(msg)
	}
}

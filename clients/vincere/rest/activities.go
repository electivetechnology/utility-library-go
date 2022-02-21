package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
)

const (
	ACTIVITY_POST_COMMENT_URL = "/activity/comment"
	COMMENT_OWNER_CANDIDATE   = "CANDIDATE"
)

type Comment struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	InsertedAt   string `json:"insert_timestamp"`
	Owner        string `json:"main_entity_type"`
	CandidateIds []int  `json:"candidate_ids"`
	CategoryIds  []int  `json:"category_ids"`
	CompanyIds   []int  `json:"company_ids"`
	PositionIds  []int  `json:"position_ids"`
	KpiActionId  int    `json:"kpi_action_id"`
}

func (client Client) CreateComment(comment Comment, token string) (Comment, error) {
	ret := Comment{}
	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), ACTIVITY_POST_COMMENT_URL, nil)

	if err != nil {
		msg := fmt.Sprintf("Error generating URL for request: %v", err)
		log.Fatalf(msg)
		return ret, errors.New(msg)
	}

	payload, err := json.Marshal(comment)
	if err != nil {
		msg := fmt.Sprintf("Could not encode Comment into Json: %v", err)
		log.Fatalf(msg)
		return ret, errors.New(msg)
	}
	reader := bytes.NewReader(payload)

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, requestUrl, reader) // URL-encoded payload

	// Add headers
	r.Header.Add("id-token", token)
	r.Header.Add("x-api-key", client.GetApiClient().GetApiKey())
	r.Header.Add("Content-Type", "application/json")

	log.Printf("Sending request with payload %v and x-api-key: %s and id-token %s", comment, client.GetApiClient().GetApiKey(), token)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		msg := fmt.Sprintf("Error Creating Comment: %v\n", err)
		log.Fatalf(msg)
		return ret, errors.New(msg)
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	switch res.StatusCode {
	case http.StatusOK:
		json.Unmarshal(data, &ret)

		return ret, nil
	default:
		msg := fmt.Sprintf("Could not create comment. Vincere returned: %v - %v", res.StatusCode, res.Body)
		log.Fatalf(msg)
		return ret, errors.New(msg)
	}
}

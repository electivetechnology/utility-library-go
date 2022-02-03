package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
)

const (
	APPLICATIONS_SEARCH_URL = "/application/filter"
)

func (client Client) SearchApplications(index int, payload *strings.Reader) error {
	// Set URL parameters on declaration
	values := url.Values{
		"index": []string{strconv.Itoa(index)},
	}

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), APPLICATIONS_SEARCH_URL, values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return errors.New("could not generate URL for application search")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPut, requestUrl, payload) // URL-encoded payload

	// Add headers
	r.Header.Add("id-token", client.GetApiClient().GetIdToken())
	r.Header.Add("x-api-key", client.GetApiClient().GetApiKey())

	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error Searching Applications: %v\n", err)
		return errors.New("error searching applications")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	return nil
}

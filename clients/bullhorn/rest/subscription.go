package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/electivetechnology/utility-library-go/clients/bullhorn"
)

const (
	REST_PUT_SUBSCRIPTION = "event/subscription/"
	REST_GET_SUBSCRIPTION = "event/subscription/"
)

func (client *Client) CreateEntitySubscription(name string, entity string, actions []string) (*bullhorn.EventsSubscription, error) {
	var entities []string

	// Add entity
	entities = append(entities, entity)

	// Forward request to multi entity/action handler
	return client.CreateSubscription(name, "entity", entities, actions)
}

func (client *Client) CreateSubscription(name string, subType string, entities []string, actions []string) (*bullhorn.EventsSubscription, error) {
	log.Printf("Will CreateSubscription subscription for the following entities %s", strings.Join(entities, ","))

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"type":        []string{"entity"},
		"names":       entities,
		"eventTypes":  []string{strings.Join(actions, ",")},
		"BhRestToken": []string{client.ApiClient.GetRestToken()},
	}

	log.Printf("Sending following data to bullhorn for new subscription: %v", values)

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), REST_PUT_SUBSCRIPTION+name, values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &bullhorn.EventsSubscription{}, errors.New("error creating new subscription")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodPut, requestUrl, nil) // URL-encoded payload

	log.Printf("Adding Header for bullhorn authorization %s", client.ApiClient.GetRestToken())

	// Print request details
	log.Printf("Request details: %v", r)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &bullhorn.EventsSubscription{}, errors.New("error creating new subscription")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		subscription := bullhorn.EventsSubscription{}
		json.Unmarshal(data, &subscription)

		// Return token
		return &subscription, nil
	}

	// If we got here there was some kind of error with exchange

	return &bullhorn.EventsSubscription{}, errors.New("error sending request to bullhorn")
}

func (client *Client) PullSubscriptionEvents(subscription string, maxEvents int) (*bullhorn.SubscriptionEvents, error) {
	log.Printf("Will pull events from bullhorn for subscription: %s", subscription)

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"maxEvents":   []string{strconv.Itoa(maxEvents)},
		"BhRestToken": []string{client.ApiClient.GetRestToken()},
	}

	log.Printf("Sending following data to bullhorn for subscription events: %v", values)

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), REST_GET_SUBSCRIPTION+subscription, values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &bullhorn.SubscriptionEvents{}, errors.New("error subscription events list")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, requestUrl, nil) // URL-encoded payload

	log.Printf("Adding Header for bullhorn authorization %s", client.ApiClient.GetRestToken())

	// Print request details
	log.Printf("Request details: %v", r)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &bullhorn.SubscriptionEvents{}, errors.New("error subscription events list")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate events
	if res.StatusCode == http.StatusOK {
		events := bullhorn.SubscriptionEvents{}
		json.Unmarshal(data, &events)

		// Return events
		return &events, nil
	}

	// If we got here there was some kind of error with exchange

	return &bullhorn.SubscriptionEvents{}, errors.New("error sending request to bullhorn")
}

func (client *Client) GetLastRequestId(subscription string) (*bullhorn.LastSubscriptionRequest, error) {
	log.Printf("Will check last request Id from bullhorn for subscription: %s", subscription)

	// Transform Authorization struct to Grant payload
	// Set URL parameters on declaration
	values := url.Values{
		"BhRestToken": []string{client.ApiClient.GetRestToken()},
	}

	log.Printf("Sending following data to bullhorn for subscription events: %v", values)

	// generate URL for request
	requestUrl, err := bullhorn.GenerateURL(client.GetApiClient().GetBaseUrl(), REST_GET_SUBSCRIPTION+subscription+"/lastRequestId", values)
	if err != nil {
		log.Printf("Error generating URL for request: %v", err)
		return &bullhorn.LastSubscriptionRequest{}, errors.New("error last request Id")
	}

	// Perform Request
	c := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, requestUrl, nil) // URL-encoded payload

	log.Printf("Adding Header for bullhorn authorization %s", client.ApiClient.GetRestToken())

	// Print request details
	log.Printf("Request details: %v", r)
	res, err := c.Do(r)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error logging in to Bullhorn rest-services: %v", err)
		return &bullhorn.LastSubscriptionRequest{}, errors.New("error subscription last request Id")
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// defer closing response body
	defer res.Body.Close()

	// print `data` as a string
	log.Printf("%s", data)

	// Success, populate lastRequest
	if res.StatusCode == http.StatusOK {
		lastRequest := bullhorn.LastSubscriptionRequest{}
		json.Unmarshal(data, &lastRequest)

		// Return lastRequest
		return &lastRequest, nil
	}

	// If we got here there was some kind of error with exchange

	return &bullhorn.LastSubscriptionRequest{}, errors.New("error sending request to bullhorn")
}

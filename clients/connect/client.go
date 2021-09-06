package connect

import (
	"log"
	"net/http"
)

type Client interface {
	IsEnabled() bool
	GetName() string
	GetId() string
}

// HandleRequest takes instance of the http.Request and performs request if client is enabled
// It returns instance of http.Response, bool (true if request was actually made, false if not)
// and error should there be any
func (client Client) HandleRequest(request *http.Request) (*http.Response, bool, error) {
	// Create new Http Client
	c := &http.Client{}

	// Set default headers
	request.Header.Set("User-Agent", client.GetName())
	request.Header.Set("X-User-Agent-Id", client.GetId())

	if !client.IsEnabled() {
		log.Printf("Client is disabled. No request will be made. Returning fake Response")
		return &http.Response{Status: "200 OK", StatusCode: 200}, false, nil
	}
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error handling request: %s %s %v", request.Method, request.URL, err)
		return res, true, err
	}

	return res, true, nil
}

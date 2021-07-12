package bullhorn

import (
	"errors"
	"net/url"
)

// Generates full URL with query string for given request
func GenerateURL(baseUrl string, path string, query url.Values) (string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Error generating URL for base: %v", err)
		return "", errors.New("error creating new url")
	}
	// Path params
	base.Path += path

	// Query params
	base.RawQuery = query.Encode()

	// Request URL
	log.Printf("Generated URL for request %q", base.String())

	return base.String(), nil
}

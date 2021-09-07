package connect

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
)

func (client Client) SaveToCache(key string, response *ApiResponse, tags []string) error {
	if client.IsCacheEnabled() {
		// Save  result to cache
		log.Printf("Saving to cache with key %s", key)

		data, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Could not encode Response into JSON %v", err)
			return err
		}

		client.Cache.SetWithTags(key, data, tags)
	}

	return nil
}

func (client Client) GetCached(request *http.Request, key string) (*ApiResponse, error) {
	if client.IsCacheEnabled() {
		// Cache is enabled check if redis is available
		log.Printf("Getting cached result for request %s: %s", request.Method, request.RequestURI)

		key := client.GetKey(key, request)
		log.Printf("Cache key for this request is: %s", key)

		data := client.Cache.Get(key)
		var res ApiResponse

		if err := json.Unmarshal(data, &res); err != nil {
			log.Printf("Failed to decode cached data %v", err)

			return &ApiResponse{}, errors.New("failed to decode cached data")
		}

		return &res, nil
	}

	log.Printf("Cache is disabled for request %s: %s", request.Method, request.RequestURI)

	return &ApiResponse{}, errors.New("response not available in cache")
}

func (client Client) GetKey(key string, request *http.Request) string {
	if key == "" {
		return client.GenerateCacheKeyFromRequest(request)
	}

	return key
}

func (client Client) GenerateKey(key string) string {
	hash := md5.Sum([]byte(key))

	return hex.EncodeToString(hash[:])
}

func (client Client) GenerateCacheKeyFromRequest(request *http.Request) string {
	h := md5.New()
	h.Write([]byte(request.Host))
	h.Write([]byte(request.Method))
	h.Write([]byte(request.RequestURI))

	hash := h.Sum(nil)

	return hex.EncodeToString(hash[:])
}

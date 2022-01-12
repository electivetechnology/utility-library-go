package config

import (
	"encoding/json"
	"github.com/electivetechnology/utility-library-go/clients/connect"
	"net/http"
)

const (
	PURPOSE_TEMPLATE_VARIABLES_URL        = "/v1/purposes/"
	PURPOSE_TEMPLATE_VARIABLES_TAG_PREFIX = "purposeTemplateVariables_"
)

type TemplateVariable struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	DisplayName  string `json:"displayName"`
	FieldFormat  string `json:"fieldFormat"`
	SampleData   string `json:"sampleData"`
	IsEnabled    string `json:"isEnabled"`
	Value        string `json:"value"`
	Organisation string `json:"organisation"`
	Visibility   string `json:"visibility"`
}

type PurposeTemplateVariable struct {
	Id                 string           `json:"id"`
	PurposeId          string           `json:"purposeId"`
	TemplateVariableId string           `json:"templateVariableId"`
	IsMultiple         string           `json:"isMultiple"`
	IsRequired         string           `json:"isRequired"`
	TemplateVariable   TemplateVariable `json:"templateVariable"`
	Purpose            Purpose          `json:"purpose"`
}

type PurposeTemplateVariableResponse struct {
	ApiResponse *connect.ApiResponse
	Data        []PurposeTemplateVariable
}

func purposeTemplateVariableRequest(path string, tagPrefix string, token string, client Client, formatData func(data []byte) []PurposeTemplateVariable) (PurposeTemplateVariableResponse, error) {
	request, _ := http.NewRequest(http.MethodGet, path, nil)
	log.Printf("purposeTemplateVariableRequest %v", path)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Option", "application/json")

	// Get key
	key := client.ApiClient.GenerateKey(tagPrefix + path + token)

	// Perform Request
	res, err := client.ApiClient.HandleRequest(request, key)

	log.Printf("purposeTemplateVariableRequest res %v", res)

	// Check for errors, purposeTemplateVariable evaluation is false
	if err != nil {
		log.Printf("Error getting Option details: %v", err)
		return PurposeTemplateVariableResponse{}, connect.NewInternalError(err.Error())
	}

	// Success, populate token
	response := PurposeTemplateVariableResponse{ApiResponse: res}

	log.Printf("purposeTemplateVariableRequest response %v", response)

	// Check if the request was actually made
	if !res.WasRequested {
		// No request was made, let's return fake response
		// This will be exactly the same token as requested for exchange
		return PurposeTemplateVariableResponse{}, nil
	}

	// read all response body
	data := res.HttpResponse.Body

	log.Printf("purposeTemplateVariableRequest data %v", data)

	// print `data` as a string
	log.Printf("%s\n", data)

	switch res.HttpResponse.StatusCode {
	case http.StatusOK:
		// Check if respose was from Cache
		if !res.WasCached {
			// Save response to cache
			log.Printf("Client provided fresh/uncached response. Saving response to cache with TTL %d", client.ApiClient.GetRedisTTL())

			// Generate tags for cache
			var tags []string
			tags = append(tags, key)

			client.ApiClient.SaveToCache(key, res, tags)
		}

		// Return response
		response.Data = formatData(data)

		return response, nil

	default:
		return response, nil
	}
}

func (client Client) GetPurposeTemplateVariables(purposeId string, token string) (PurposeTemplateVariableResponse, error) {
	log.Printf("Will request purposeTemplateVariable option with purposeTemplateVariableId")

	path := client.ApiClient.GetBaseUrl() + PURPOSE_TEMPLATE_VARIABLES_URL + purposeId + "/template-variables"

	var formatData = func(data []byte) []PurposeTemplateVariable {
		var responseData []PurposeTemplateVariable
		json.Unmarshal(data, &responseData)
		return responseData
	}

	return purposeTemplateVariableRequest(path, PURPOSE_TEMPLATE_VARIABLES_TAG_PREFIX, token, client, formatData)
}

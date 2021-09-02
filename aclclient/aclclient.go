package aclclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/electivetechnology/utility-library-go/auth"
	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/gin-gonic/gin"
)

const (
	AUTH_URL           = "/v1/authorise"
	TOKEN_EXCHANGE_URL = "/v1/token/exchange"
)

type AclClient struct {
	AclHost string
}

type AclCheck struct {
	Name         string `json:"name"`
	Subject      string `json:"subject"`
	Permission   string `json:"permission"`
	Organisation string `json:"organisation"`
}

type AuthorisedResponse struct {
	Message string
	Checks  []string
}

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("aclclient")
}

func NewAclCheck(subject string, permission string) *AclCheck {
	check := &AclCheck{Name: "main"}
	check.Subject = subject
	check.Permission = permission

	return check
}

func NewAclClient() *AclClient {
	// Get Host
	host := os.Getenv("ACL_HOST")

	if host == "" {
		host = "http://localhost:8011"
	}

	client := &AclClient{AclHost: host}

	return client
}

func (client AclClient) IsTokenAuthorised(token string, aclCheck *AclCheck) bool {
	// Create new Http Client
	c := &http.Client{}

	// Transform AclCheck struct to json payload
	jsonValue, _ := json.Marshal(aclCheck)
	request, _ := http.NewRequest(http.MethodPost, client.AclHost+TOKEN_EXCHANGE_URL, bytes.NewBuffer(jsonValue))
	log.Printf("Checking if user have %s permissions on subject %s", aclCheck.Permission, aclCheck.Subject)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")

	// Perform Request
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error processing Authorisation: %v\n", err)
		return false
	}

	if res.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func (client AclClient) IsRequestAuthorized(ctx *gin.Context, aclCheck *AclCheck) bool {
	// Get SecurityToken
	st, _ := ctx.Get("SecurityToken")
	token := st.(auth.SecurityToken)
	aclCheck.Organisation = token.GetOrganisation()

	return client.IsTokenAuthorised(token.GetRawToken(), aclCheck)
}

func (client AclClient) ExchangeToken(payload *ExchangePaylod) (ExchangeResponse, error) {
	// Create new Http Client
	c := &http.Client{}

	// Transform Token struct to json payload
	jsonValue, _ := json.Marshal(payload)
	request, _ := http.NewRequest(http.MethodPost, client.AclHost+AUTH_URL, bytes.NewBuffer(jsonValue))
	log.Printf("Exchanging token for organisation %s", payload.Organisation)

	// Set Headers for this request
	request.Header.Set("Authorization", "Bearer "+payload.Token)
	request.Header.Add("Content-Type", "application/json")

	// Perform Request
	res, err := c.Do(request)

	// Check for errors, default evaluation is false
	if err != nil {
		log.Printf("Error processing Token exchange: %v\n", err)
		return ExchangeResponse{}, err
	}

	// read all response body
	data, _ := ioutil.ReadAll(res.Body)

	// print `data` as a string
	log.Printf("%s\n", data)

	// Success, populate token
	if res.StatusCode == http.StatusOK {
		result := ExchangeResponse{}
		json.Unmarshal(data, &result)

		// Return token
		return result, nil
	}

	return ExchangeResponse{}, errors.New("error exchanging Token for new organisation")
}

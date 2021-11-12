package acl

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/electivetechnology/utility-library-go/auth"
	"github.com/gin-gonic/gin"
)

const (
	AUTH_URL = "/v1/authorise"
)

type Authorise struct {
	Subject      string `json:"subject"`
	Permission   string `json:"permission"`
	Organisation string `json:"organisation"`
}

type Checks struct {
	Name      string    `json:"name"`
	Authorise Authorise `json:"authorise"`
}

type AclCheck struct {
	Name         string   `json:"name"`
	Subject      string   `json:"subject"`
	Permission   string   `json:"permission"`
	Organisation string   `json:"organisation"`
	Checks       []Checks `json:"checks"`
}

func NewAclCheck(subject string, permission string) *AclCheck {
	check := &AclCheck{Name: "main"}
	check.Subject = subject
	check.Permission = permission

	return check
}

func AddAclCheck(ctx *gin.Context, aclCheck *AclCheck, name string, subject string, permission string) *AclCheck {
	// Get SecurityToken
	st, _ := ctx.Get("SecurityToken")
	token := st.(auth.SecurityToken)

	aclCheck.Checks = []Checks{{
		Name: name,
		Authorise: Authorise{
			Permission:   permission,
			Subject:      subject,
			Organisation: token.GetOrganisation(),
		},
	}}

	return aclCheck
}

func (client Client) IsTokenAuthorised(token string, aclCheck *AclCheck) bool {
	if !client.ApiClient.IsEnabled() {
		return true
	}
	// Create new Http Client
	c := &http.Client{}

	// Transform AclCheck struct to json payload
	jsonValue, _ := json.Marshal(aclCheck)
	request, _ := http.NewRequest(http.MethodPost, client.ApiClient.GetBaseUrl()+AUTH_URL, bytes.NewBuffer(jsonValue))
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

func (client Client) IsRequestAuthorized(ctx *gin.Context, aclCheck *AclCheck) bool {
	// Get SecurityToken
	st, _ := ctx.Get("SecurityToken")
	token := st.(auth.SecurityToken)
	aclCheck.Organisation = token.GetOrganisation()

	return client.IsTokenAuthorised(token.GetRawToken(), aclCheck)
}

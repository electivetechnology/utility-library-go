package rest

import "net/http"

type ApiClient interface {
	GetBaseUrl() string
	GetIdToken() string
	SetIdToken(token string)
	GetHttpClient() *http.Client
	SetHttpClient(h *http.Client)
	GetApiKey() string
}

type BaseClient struct {
	BaseUrl    string
	IdToken    string
	ApiKey     string
	HttpClient *http.Client
}

func (c BaseClient) GetBaseUrl() string {
	return c.BaseUrl
}

func (c BaseClient) GetIdToken() string {
	return c.IdToken
}

func (c *BaseClient) SetIdToken(token string) {
	c.IdToken = token
}

func (c BaseClient) GetHttpClient() *http.Client {
	return c.HttpClient
}

func (c *BaseClient) SetHttpClient(h *http.Client) {
	c.HttpClient = h
}

func (c BaseClient) GetApiKey() string {
	return c.ApiKey
}

package rest

import (
	"time"
)

type RestToken struct {
	BhRestToken string `json:"BhRestToken"`
	RestUrl     string `json:"restUrl"`
	Ttl         int
	ExpiresAt   *time.Time
}

func (client *Client) GetBhRestToken(accessToken string) (*RestToken, error) {
	return client.Login(accessToken)
}

func (client *Client) AddRestToken(token *RestToken) {
	client.GetApiClient().SetBaseUrl(token.RestUrl)
	client.GetApiClient().SetRestToken(token.BhRestToken)
}

package oauth

import "time"

type Token interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresIn() int
	GetRefreshToken() string
	GetExpiresAt() *time.Time
	GetIdToken() string
}

type AccessToken struct {
	AccessToken string     `json:"access_token"`
	TokenType   string     `json:"token_type"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

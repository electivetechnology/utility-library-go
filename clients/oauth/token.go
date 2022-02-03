package oauth

import (
	"time"

	"github.com/electivetechnology/utility-library-go/clients/connect"
)

type Token interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresIn() int
	GetRefreshToken() string
	GetExpiresAt() *time.Time
	GetIdToken() string
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    *time.Time
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

type Response struct {
	ApiResponse *connect.ApiResponse
	Token       Token
}

type TokenResponse interface {
	GetToken() Token
	GetApiResponse() *connect.ApiResponse
}

func (r Response) GetToken() Token {
	return r.Token
}

func (r Response) GetApiResponse() *connect.ApiResponse {
	return r.ApiResponse
}

func (t AccessToken) GetAccessToken() string {
	return t.AccessToken
}

func (t AccessToken) GetTokenType() string {
	return t.TokenType
}

func (t AccessToken) GetExpiresIn() int {
	return t.ExpiresIn
}

func (t AccessToken) GetRefreshToken() string {
	return t.RefreshToken
}

func (t AccessToken) GetExpiresAt() *time.Time {
	return t.ExpiresAt
}

func (t AccessToken) GetIdToken() string {
	return t.IdToken
}

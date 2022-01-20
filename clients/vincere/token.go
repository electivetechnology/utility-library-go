package vincere

import "time"

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    *time.Time
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

func (t Token) GetAccessToken() string {
	return t.AccessToken
}

func (t Token) GetExpiresAt() *time.Time {
	return t.ExpiresAt
}

func (t Token) GetExpiresIn() int {
	return t.ExpiresIn
}

func (t Token) GetRefreshToken() string {
	return t.RefreshToken
}

func (t Token) GetTokenType() string {
	return t.TokenType
}

func (t Token) GetIdToken() string {
	return t.IdToken
}

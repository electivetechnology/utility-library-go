package oauth

import "time"

type Token interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresIn() int
	GetRefreshToken() string
	GetExpiresAt() *time.Time
}

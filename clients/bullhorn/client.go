package bullhorn

import "os"

const (
	AUTH_TOKEN_URL = "/oauth/token"
)

type OAuthClient struct {
	BaseUrl string
}

func NewOAuthClient() *OAuthClient {
	// Get Base URL
	url := os.Getenv("BULLHORN_OAUTH_BASE_URL")

	if url == "" {
		url = "https://auth.bullhornstaffing.com"
	}

	return &OAuthClient{BaseUrl: url}
}

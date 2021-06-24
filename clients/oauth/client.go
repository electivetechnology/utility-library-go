package oauth

type OAuthClient interface {
	GetToken(auth Authorization) (Token, error)
	RefreshToken(Token Token, clientId string, clientSecret string) (Token, error)
	Refresh(refreshToken string) (Token, error)
}

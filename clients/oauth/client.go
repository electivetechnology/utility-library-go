package oauth

type OAuthClient interface {
	GetToken(auth Authorization) (Token, error)
}

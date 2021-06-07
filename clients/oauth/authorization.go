package oauth

type Authorization interface {
	GetCode() string
	GetClientId() string
	GetClientSecret() string
}

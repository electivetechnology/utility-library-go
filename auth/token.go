package auth

type Token struct {
	RawToken      string
	Username      string `json:"username"`
	Organisation  string `json:"organisation"`
	Authenticated bool
}

type SecurityToken interface {
	IsAuthenticated() bool
	GetOrganisation() string
	GetRawToken() string
}

func NewToken() Token {
	return Token{
		Authenticated: false,
	}
}

func (t Token) IsAuthenticated() bool {
	if t.Authenticated {
		return true
	}

	return false
}

func (t Token) GetOrganisation() string {
	return t.Organisation
}

func (t Token) GetUsername() string {
	return t.Username
}

func (t Token) GetRawToken() string {
	return t.RawToken
}

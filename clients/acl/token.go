package acl

type ExchangePayload struct {
	Organisation string `json:"organisation"`
	Token        string `json:"token"`
}

type ExchangeResponse struct {
	Token string `json:"token"`
}

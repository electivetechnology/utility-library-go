package aclclient

type ExchangePaylod struct {
	Organisation string `json:"organisation"`
	Token        string `json:"organisation"`
}

type ExchangeResponse struct {
	Token string `json:"token"`
}

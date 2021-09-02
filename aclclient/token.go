package aclclient

type ExchangePaylod struct {
	Organisation string `json:"organisation"`
	Token        string `json:"token"`
}

type ExchangeResponse struct {
	Token string `json:"token"`
}

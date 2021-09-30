package components

type ComonentNotification struct {
	ID       string `json:"id"`
	Callback string `json:"callback"`
}

type ComponentStatusResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

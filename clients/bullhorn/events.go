package bullhorn

type EventsSubscription struct {
	Id            string `json:"subscriptionId"`
	LastRequestId int    `json:"lastRequestId"`
	CreatedOn     int    `json:"createdOn"`
}

package bullhorn

type EventsSubscription struct {
	Id            string `json:"subscriptionId"`
	LastRequestId int    `json:"lastRequestId"`
	CreatedOn     int    `json:"createdOn"`
}

type SubscriptionEvents struct {
	RequestId int     `json:"requestId"`
	EventList []Event `json:"events"`
}

type Event struct {
	Id                string   `json:"eventId"`
	Type              string   `json:"eventType"`
	TimeStamp         int      `json:"eventTimestamp"`
	EntityName        string   `json:"entityName"`
	EntityId          int      `json:"entityId"`
	EntityEventType   string   `json:"entityEventType"`
	UpdatedProperties []string `json:"updatedProperties"`
}

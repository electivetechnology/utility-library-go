package contracts

import "time"

type MessageInterface interface {
	GetTo() string
	GetFrom() string
	GetSubject() string
	GetContent() string
	GetSendAfter() time.Time
}

type Message struct {
	To        string    `json:"to"`
	From      string    `json:"from"`
	Content   string    `json:"content"`
	Subject   string    `json:"subject"`
	SendAfter time.Time `json:"send_after"`
}

func (message Message) GetTo() string {
	return message.To
}

func (message Message) GetFrom() string {
	return message.From
}

func (message Message) GetSubject() string {
	return message.Subject
}

func (message Message) GetContent() string {
	return message.Content
}

func (message Message) GetSendAfter() time.Time {
	return message.SendAfter
}

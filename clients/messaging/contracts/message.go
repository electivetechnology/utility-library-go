package contracts

import "time"

type MessageInterface interface {
	GetTo() string
	GetFrom() string
	GetSubject() string
	GetContent() string
	GetSendAfter() time.Time
}

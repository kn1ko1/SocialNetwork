package models

type Message struct {
	MessageId   int
	Body        string
	CreatedAt   int64
	MessageType string
	SenderId    int
	TargetId    int
	UpdatedAt   int64
}

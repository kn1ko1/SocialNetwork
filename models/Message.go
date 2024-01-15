package models

type Message struct {
	MessageId   int
	Body        string
	CreatedAt   int
	MessageType string
	SenderId    int
	TargetId    int
	UpdatedAt   int64
}

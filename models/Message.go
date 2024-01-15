package models

type Message struct {
	MessageId int
	Body      string
	CreatedAt int64
	SenderId  int
	TargetId  int
	Type      string
	UpdatedAt int64
}

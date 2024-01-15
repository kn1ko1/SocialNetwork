package models

type Message struct {
	MessageId int
	Body string
	CreatedAt int
	SenderId int
	TargetId int
	Type string
	UpdatedAt int64
}
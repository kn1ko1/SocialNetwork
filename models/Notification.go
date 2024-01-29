package models

type Notification struct {
	NotificationId   int
	CreatedAt        int64
	NotificationType string
	ObjectId         int64
	SenderId         int64
	Status           string
	TargetId         int
	UpdatedAt        int64
}

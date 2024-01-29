package models

import (
	"errors"
	"math/rand"
)

type Notification struct {
	NotificationId   int
	CreatedAt        int64
	NotificationType string
	ObjectId         int
	SenderId         int
	Status           string
	TargetId         int
	UpdatedAt        int64
}

func (n *Notification) Validate() error {
	if n.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if n.NotificationType == "" {
		return errors.New("Notification Type must not be empty")
	}
	if n.SenderId <= 0 {
		return errors.New("invalid 'SenderId' field")
	}
	if n.Status == "" {
		return errors.New("notification status must not be empty")
	}
	if n.TargetId <= 0 {
		return errors.New("invalid 'PostId' field")
	}
	if n.UpdatedAt < n.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	return nil
}

func GenerateValidNotification() *Notification {
	ctime := rand.Int63n(1000) + 1
	idxNotificationType := rand.Intn(len(sutNotificationType))
	idxNotificationStatus := rand.Intn(len(sutNotificationStatus))
	n := &Notification{
		CreatedAt:        ctime,
		NotificationType: sutNotificationType[idxNotificationType],
		ObjectId:         rand.Intn(1000) + 1,
		SenderId:         rand.Intn(1000) + 1,
		Status:           sutNotificationStatus[idxNotificationStatus],
		TargetId:         rand.Intn(1000) + 1,
		UpdatedAt:        ctime,
	}
	return n
}

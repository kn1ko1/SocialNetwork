package models

import (
	"errors"
	"math/rand"
)

type Notification struct {
	NotificationId   int    `json:"notificationId"`
	CreatedAt        int64  `json:"createdAt"`
	NotificationType string `json:"notificationType"`
	ObjectId         int    `json:"objectId"`
	SenderId         int    `json:"senderId"`
	Status           string `json:"status"`
	TargetId         int    `json:"targetId"`
	UpdatedAt        int64  `json:"updatedAt"`
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
		return errors.New("invalid 'TargetId' field")
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

func GenerateMissingFieldNotification() *Notification {
	n := GenerateValidNotification()
	missingField := rand.Intn(6)
	switch missingField {
	case 0:
		n.CreatedAt = 0
	case 1:
		n.NotificationType = ""
	case 2:
		n.SenderId = 0
	case 3:
		n.Status = ""
	case 4:
		n.TargetId = 0
	case 5:
		n.UpdatedAt = 0
	}
	return n
}

func GenerateInvalidNotification() *Notification {
	n := GenerateValidNotification()
	invalidField := rand.Intn(6)
	switch invalidField {
	case 0:
		n.CreatedAt = -n.CreatedAt
	case 1:
		n.NotificationType = ""
	case 2:
		n.SenderId = -n.SenderId
	case 3:
		n.Status = ""
	case 4:
		n.TargetId = -n.TargetId
	case 5:
		n.UpdatedAt = -n.UpdatedAt
	}
	return n
}

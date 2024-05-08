package transport

import "socialnetwork/models"

type NotificationResponse struct {
	Reply        string              `json:"reply"`
	Notification models.Notification `json:"notification"`
}

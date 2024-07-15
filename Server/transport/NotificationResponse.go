package transport

import "socialnetwork/Server/models"

type NotificationResponse struct {
	Reply        string              `json:"reply"`
	Notification models.Notification `json:"notification"`
}

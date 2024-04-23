package repo

import "socialnetwork/models"

type INotifications interface {
	//Notification
	CreateNotification(notification models.Notification) (models.Notification, error)
	GetNotificationById(notificationId int) (models.Notification, error)
	GetNotificationsByTargetId(targetId int) ([]models.Notification, error)
	UpdateNotification(notification models.Notification) (models.Notification, error)
	DeleteNotificationById(notificationId int) error
}

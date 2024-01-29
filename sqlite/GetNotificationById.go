package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves notification with the relevant notificationId from the NOTIFICATION table
func GetNotificationById(database *sql.DB, notificationId int) (models.Notification, error) {
	var notification models.Notification
	err := database.QueryRow("SELECT * FROM NOTIFICATIONS WHERE NotificationId = ?", notificationId).
		Scan(
			&notification.CreatedAt,
			&notification.NotificationType,
			&notification.ObjectId,
			&notification.SenderId,
			&notification.Status,
			&notification.TargetId,
			&notification.UpdatedAt,
		)

	switch {
	case err == sql.ErrNoRows:
		return notification, errors.New("notification not found")
	case err != nil:
		return notification, err
	}

	return notification, nil
}

package notifications

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Retrieves notification with the relevant notificationId from the NOTIFICATIONS table
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
		utils.HandleError("Notification not found.", err)
		return notification, errors.New("notification not found")
	case err != nil:
		utils.HandleError("Error querying notification by ID.", err)
		return notification, err
	}

	return notification, nil
}

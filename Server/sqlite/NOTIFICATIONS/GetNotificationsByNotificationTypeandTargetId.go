package notifications

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves notifications with the relevant userId from the NOTIFICATIONS table
func GetNotificationsByTargetId(database *sql.DB, targetId int) ([]models.Notification, error) {
	var notifications []models.Notification

	rows, err := database.Query("SELECT * FROM NOTIFICATIONS WHERE TargetId = ?", targetId)
	if err != nil {
		return notifications, nil
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(
			&notification.NotificationId,
			&notification.CreatedAt,
			&notification.NotificationType,
			&notification.ObjectId,
			&notification.SenderId,
			&notification.Status,
			&notification.TargetId,
			&notification.UpdatedAt,
		)
		if err != nil {
			utils.HandleError("Error scanning row in GetNotificationsByUserId.", err)
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetNotificationsByUserId.", err)
		return nil, err
	}

	return notifications, nil
}

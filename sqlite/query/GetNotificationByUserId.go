package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves notifications with the relevant userId from the NOTIFICATIONS table
func GetNotificationsByUserId(database *sql.DB, userId int) ([]models.Notification, error) {
	rows, err := database.Query("SELECT * FROM NOTIFICATIONS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error querying notifications by UserId.", err)
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(
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

package notifications

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Updates notification information in the MESSAGES table
func UpdateNotification(database *sql.DB, notification models.Notification) (models.Notification, error) {
	query := `
		UPDATE NOTIFICATIONS
		SET
			NotificationType = ?,
			Status = ?,
			UpdatedAt = ?
		WHERE NotificationId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing database in UpdateNotification.", err)
		return notification, err
	}

	_, err = statement.Exec(
		notification.NotificationType,
		notification.Status,
		notification.UpdatedAt,
		notification.NotificationId,
	)

	if err != nil {
		utils.HandleError("Error executing statement in UpdateNotification.", err)
		return notification, err
	}

	return notification, nil
}

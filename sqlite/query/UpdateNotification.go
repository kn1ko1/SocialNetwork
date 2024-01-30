package sqlite

import (
	"database/sql"
	"socialnetwork/models"
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
		return notification, err
	}

	_, err = statement.Exec(
		&notification.NotificationType,
		&notification.Status,
		&notification.UpdatedAt,
		&notification.NotificationId,
	)

	if err != nil {
		return notification, err
	}

	return notification, nil
}

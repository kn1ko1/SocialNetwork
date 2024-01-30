package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds notification into the given database
func CreateNotification(database *sql.DB, notification models.Notification) (models.Notification, error) {

	query := `
	INSERT INTO MESSAGES (
		CreatedAt,
		NotificationType,
		ObjectId,
		SenderId,
		Status,
		TargetId,
		UpdatedAt
	) VALUES (?, ?, ?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		return notification, err
	}
	res, err := statement.Exec(
		notification.CreatedAt,
		notification.NotificationType,
		notification.ObjectId,
		notification.SenderId,
		notification.Status,
		notification.TargetId,
		notification.UpdatedAt,
	)

	if err != nil {
		return notification, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return notification, err
	}
	notification.NotificationId = int(id)
	return notification, nil
}

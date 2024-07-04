package notifications

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Adds notification into the given database
func CreateNotification(database *sql.DB, notification models.Notification) (models.Notification, error) {

	// Check if a notification with the same NotificationType, ObjectId, SenderId, and TargetId already exists
	existingNotification := models.Notification{}
	query := `
		SELECT NotificationId, CreatedAt, NotificationType, ObjectId, SenderId, Status, TargetId, UpdatedAt
		FROM NOTIFICATIONS
		WHERE NotificationType = ? AND ObjectId = ? AND SenderId = ? AND TargetId = ?
	`

	err := database.QueryRow(query, notification.NotificationType, notification.ObjectId, notification.SenderId, notification.TargetId).Scan(
		&existingNotification.NotificationId,
		&existingNotification.CreatedAt,
		&existingNotification.NotificationType,
		&existingNotification.ObjectId,
		&existingNotification.SenderId,
		&existingNotification.Status,
		&existingNotification.TargetId,
		&existingNotification.UpdatedAt,
	)

	if err == nil {
		return existingNotification, errors.New("notification already exists in db")
	} else if err != sql.ErrNoRows {
		utils.HandleError("Error querying db for existing notification.", err)
		return notification, err
	}

	// Insert the new notification into the database
	insertQuery := `
		INSERT INTO NOTIFICATIONS (
			CreatedAt,
			NotificationType,
			ObjectId,
			SenderId,
			Status,
			TargetId,
			UpdatedAt
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	statement, err := database.Prepare(insertQuery)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
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
		utils.HandleError("Error executing statement.", err)
		return notification, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return notification, err
	}
	notification.NotificationId = int(id)
	return notification, nil
}

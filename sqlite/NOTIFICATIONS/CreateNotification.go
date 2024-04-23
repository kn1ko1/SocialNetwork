package notifications

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Adds notification into the given database
func CreateNotification(database *sql.DB, notification models.Notification) (models.Notification, error) {

	query := `
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

	statement, err := database.Prepare(query)
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

package notifications

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes a specific user from the USERS table
func DeleteNotificationById(database *sql.DB, notificationId int) error {
	_, err := database.Exec("DELETE FROM NOTIFICATIONS WHERE NotificationId = ?", notificationId)
	if err != nil {
		utils.HandleError("Error executing delete notifications by NotificationId statement.", err)
		return err
	}
	return nil
}

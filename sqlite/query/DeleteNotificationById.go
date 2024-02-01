package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes a specific user from the USERS table
func DeleteNotificationById(database *sql.DB, notificationId int) error {
	_, err := database.Exec("DELETE FROM NOTIFICATION WHERE NotificationId = ?", notificationId)
	if err != nil {
		utils.HandleError("Error executing delete notifications by NotificationId statement.", err)
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes a specific user from the USERS table
func DeleteNotificationById(database *sql.DB, notificationId int) error {
	_, err := database.Exec("DELETE FROM NOTIFICATION WHERE NotificationId = ?", notificationId)
	if err != nil {
		return err
	}
	return nil
}

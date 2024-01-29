package sqlite

import "database/sql"

// deletes event users related to UserId from the EVENT_USERS table
func DeleteEventUsersByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM EVENT_USERS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

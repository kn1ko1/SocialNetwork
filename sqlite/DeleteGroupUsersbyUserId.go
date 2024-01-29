package sqlite

import "database/sql"

// deletes group users related to UserId from the GROUP_USERS table
func DeleteGroupUsersByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

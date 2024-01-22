package sqlite

import "database/sql"

// deletes a specific user from the USERS table
func DeleteUserById(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM USERS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

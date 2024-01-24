package sqlite

import "database/sql"

// Deletes all users from the USERS table; use with caution
func DeleteAllUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM USERS")
	if err != nil {
		return err
	}
	return nil
}

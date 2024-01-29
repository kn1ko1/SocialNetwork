package sqlite

import "database/sql"

// Deletes all eventusers from the EVENT_USERS table; use with caution
func DeleteAllEventUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM EVENT_USERS")
	if err != nil {
		return err
	}
	return nil
}

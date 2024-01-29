package sqlite

import "database/sql"

// Deletes all groupusers from the GROUP_USERS table; use with caution
func DeleteAllGroupUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS")
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// Deletes all comments from the COMMENTS table; use with caution
func DeleteAllGroups(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM GROUPS")
	if err != nil {
		return err
	}
	return nil
}

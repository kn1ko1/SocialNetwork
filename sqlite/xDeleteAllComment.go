package sqlite

import "database/sql"

// Deletes all comments from the COMMENTS table; use with caution
func DeleteAllComments(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM COMMENTS")
	if err != nil {
		return err
	}
	return nil
}

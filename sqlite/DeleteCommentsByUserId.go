package sqlite

import "database/sql"

// deletes comments related to UserId from the COMMENTS table
func DeleteCommentsByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes comments related to UserId from the COMMENTS table
func DeleteCommentsByUserId(db *sql.DB, userId int) error {
	_, err := db.Exec("DELETE FROM COMMENTS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes comments related to CommentId from the COMMENTS table
func DeleteCommentsById(database *sql.DB, commentId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE CommentId = ?", commentId)
	if err != nil {
		return err
	}
	return nil
}

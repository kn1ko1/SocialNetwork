package sqlite

import "database/sql"

// deletes posts related to PostId from the COMMENTS table
func DeleteCommentsByPostId(database *sql.DB, postId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE PostId = ?", postId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes posts related to PostId from the COMMENTS table
func DeleteCommentsByPostId(db *sql.DB, postId int) error {
	_, err := db.Exec("DELETE FROM COMMENTS WHERE PostId = ?", postId)
	if err != nil {
		return err
	}
	return nil
}

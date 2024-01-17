package sqlite

import "database/sql"

// deletes a specific post from the POSTS table
func DeletePostById(db *sql.DB, postId int) error {
	_, err := db.Exec("DELETE FROM POSTS WHERE PostId = ?", postId)
	if err != nil {
		return err
	}
	return nil
}

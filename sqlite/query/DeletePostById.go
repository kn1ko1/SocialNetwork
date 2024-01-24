package sqlite

import "database/sql"

// deletes a specific post from the POSTS table
func DeletePostById(database *sql.DB, postId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE PostId = ?", postId)
	if err != nil {
		return err
	}
	return nil
}

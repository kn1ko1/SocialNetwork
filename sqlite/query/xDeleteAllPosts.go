package sqlite

import "database/sql"

// Deletes all posts from the POSTS table; use with caution
func DeleteAllPosts(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM POSTS")
	if err != nil {
		return err
	}
	return nil
}

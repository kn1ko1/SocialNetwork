package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all posts from the POSTS table; use with caution
func DeleteAllPosts(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM POSTS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

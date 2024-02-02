package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes posts related to postId from the POST_USERS table
func DeletePostUsersByPostId(database *sql.DB, postId int) error {
	_, err := database.Exec("DELETE FROM POST_USERS WHERE postId = ?", postId)
	if err != nil {
		utils.HandleError("Error executing delete post users by post ID statement.", err)
		return err
	}
	return nil
}

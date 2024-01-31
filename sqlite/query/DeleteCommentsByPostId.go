package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes posts related to PostId from the COMMENTS table
func DeleteCommentsByPostId(database *sql.DB, postId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE PostId = ?", postId)
	if err != nil {
		utils.HandleError("Error executing delete comments by post ID statement.", err)
		return err
	}
	return nil
}

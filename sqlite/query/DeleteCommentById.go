package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes comments related to CommentId from the COMMENTS table
func DeleteCommentById(database *sql.DB, commentId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE CommentId = ?", commentId)
	if err != nil {
		utils.HandleError("Error executing delete comment statement.", err)
		return err
	}
	return nil
}

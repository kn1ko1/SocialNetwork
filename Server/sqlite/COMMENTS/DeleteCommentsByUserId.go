package comments

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// deletes comments related to UserId from the COMMENTS table
func DeleteCommentsByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete comments by user ID statement.", err)
		return err
	}
	return nil
}

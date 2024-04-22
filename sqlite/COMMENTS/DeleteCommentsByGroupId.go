package comments

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes comments related to groupId from the COMMENTS table
func DeleteCommentsByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM COMMENTS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete comments by group ID statement.", err)
		return err
	}
	return nil
}

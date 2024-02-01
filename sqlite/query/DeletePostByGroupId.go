package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes posts related to groupId from the POSTS table
func DeletePostByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE groupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete posts by groupId statement.", err)
		return err
	}
	return nil
}

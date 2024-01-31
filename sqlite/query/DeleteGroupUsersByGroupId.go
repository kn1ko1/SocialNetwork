package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes posts related to groupId from the GROUP_USERS table
func DeleteGroupUserByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE groupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete group users by group Id.", err)
		return err
	}
	return nil
}

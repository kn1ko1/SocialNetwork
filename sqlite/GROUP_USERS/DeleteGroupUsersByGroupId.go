package group_users

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes posts related to GroupId from the GROUP_USERS table
func DeleteGroupUserByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete group users by group Id.", err)
		return err
	}
	return nil
}

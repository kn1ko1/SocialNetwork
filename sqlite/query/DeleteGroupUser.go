package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes groupo of GroupId from the GROUPS table
func DeleteGroupUser(database *sql.DB, groupUserId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE GroupUserId = ?", groupUserId)
	if err != nil {
		utils.HandleError("Error executing delete group users by group user ID statement.", err)
		return err
	}
	return nil
}

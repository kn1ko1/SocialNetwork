package group_users

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// deletes group users related to UserId from the GROUP_USERS table
func DeleteGroupUsersByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete group users by user ID statement.", err)
		return err
	}
	return nil
}

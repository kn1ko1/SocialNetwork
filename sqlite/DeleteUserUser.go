package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes groupo of GroupId from the GROUPS table
func DeleteUserUser(database *sql.DB, userUserId int) error {
	_, err := database.Exec("DELETE FROM USER_USERS WHERE UserUserId = ?", userUserId)
	if err != nil {
		utils.HandleError("Error executing delete user users by user user ID statement.", err)
		return err
	}
	return nil
}

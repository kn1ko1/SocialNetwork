package group_users

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// Deletes all groupusers from the GROUP_USERS table; use with caution
func DeleteAllGroupUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

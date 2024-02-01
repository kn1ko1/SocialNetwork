package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
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

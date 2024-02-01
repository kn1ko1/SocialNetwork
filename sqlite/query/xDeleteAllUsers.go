package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all users from the USERS table; use with caution
func DeleteAllUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM USERS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

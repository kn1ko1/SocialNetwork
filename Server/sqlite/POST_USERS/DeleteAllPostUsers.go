package post_users

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// Deletes all postusers from the POST_USERS table; use with caution
func DeleteAllPostUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM POST_USERS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

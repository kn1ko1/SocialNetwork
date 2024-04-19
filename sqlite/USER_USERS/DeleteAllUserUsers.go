package user_users

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all userusers from the USER_USERS table; use with caution
func DeleteAllUserUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM USER_USERS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all eventusers from the EVENT_USERS table; use with caution
func DeleteAllEventUsers(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM EVENT_USERS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

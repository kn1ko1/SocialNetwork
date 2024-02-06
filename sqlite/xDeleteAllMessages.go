package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all messages from the MESSAGES table; use with caution
func DeleteAllMessages(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM MESSAGES")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes message of MessageId from the MESSAGES table
func DeleteMessageById(database *sql.DB, messageId int) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE MessageId = ?", messageId)
	if err != nil {
		utils.HandleError("Error executing delete messages by message Id statement.", err)
		return err
	}
	return nil
}

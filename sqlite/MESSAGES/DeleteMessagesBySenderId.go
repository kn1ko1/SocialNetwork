package messages

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes messages by SenderId from the MESSAGES table
func DeleteMessagesBySenderId(database *sql.DB, senderId int) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE SenderId = ?", senderId)
	if err != nil {
		utils.HandleError("Error executing delete messages by SenderId statement.", err)
		return err
	}
	return nil
}

package messages

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// deletes messages of MessageType from the MESSAGES table
func DeleteMessagesByType(database *sql.DB, messageType string) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE MessageType = ?", messageType)
	if err != nil {
		utils.HandleError("Error executing delete messages by MessageType statement.", err)
		return err
	}
	return nil
}

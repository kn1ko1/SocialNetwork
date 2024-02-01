package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes messages by TargetId from the MESSAGES table
func DeleteMessagesByTargetId(database *sql.DB, targetId int) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE TargetId = ?", targetId)
	if err != nil {
		utils.HandleError("Error executing delete messages by TargetId statement.", err)
		return err
	}
	return nil
}

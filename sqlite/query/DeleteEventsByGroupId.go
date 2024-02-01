package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes events related to GroupId from the EVENTS table
func DeleteEventsByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM EVENTS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete events by group ID statement.", err)
		return err
	}
	return nil
}

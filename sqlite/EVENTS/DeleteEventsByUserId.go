package events

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes events related to UserId from the EVENTS table
func DeleteEventsByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM EVENTS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete events by user ID statement.", err)
		return err
	}
	return nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes event related to EventId from the EVENTS table
func DeleteEventById(database *sql.DB, eventId int) error {
	_, err := database.Exec("DELETE FROM EVENTS WHERE EventId = ?", eventId)
	if err != nil {
		utils.HandleError("Error executing delete event by ID statement.", err)
		return err
	}
	return nil
}

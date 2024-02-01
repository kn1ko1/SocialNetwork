package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// Deletes all events from the EVENTS table; use with caution
func DeleteAllEvents(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM EVENTS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

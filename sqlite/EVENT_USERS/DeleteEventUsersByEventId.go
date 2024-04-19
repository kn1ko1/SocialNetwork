package event_users

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes posts related to groupId from the EVENT_USERS table
func DeleteEventUsersByEventId(database *sql.DB, eventId int) error {
	_, err := database.Exec("DELETE FROM EVENT_USERS WHERE eventId = ?", eventId)
	if err != nil {
		utils.HandleError("Error executing delete event users by event ID statement.", err)
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes posts related to groupId from the EVENT_USERS table
func DeleteEventUserByEventId(database *sql.DB, eventId int) error {
	_, err := database.Exec("DELETE FROM EVENT_USERS WHERE eventId = ?", eventId)
	if err != nil {
		return err
	}
	return nil
}

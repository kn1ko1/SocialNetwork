package sqlite

import "database/sql"

// deletes event related to EventId from the EVENTS table
func DeleteEventById(db *sql.DB, eventId int) error {
	_, err := db.Exec("DELETE FROM EVENTS WHERE EventId = ?", eventId)
	if err != nil {
		return err
	}
	return nil
}

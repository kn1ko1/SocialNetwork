package sqlite

import "database/sql"

// Deletes all events from the EVENTS table; use with caution
func DeleteAllEvents(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM EVENTS")
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes events related to UserId from the EVENT table
func DeleteEventsByUserId(db *sql.DB, userId int) error {
	_, err := db.Exec("DELETE FROM EVENTS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

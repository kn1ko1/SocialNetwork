package sqlite

import "database/sql"

// deletes messages by SenderId from the MESSAGES table
func DeleteMessagesBySenderId(db *sql.DB, senderId string) error {
	_, err := db.Exec("DELETE FROM MESSAGES WHERE SenderId = ?", senderId)
	if err != nil {
		return err
	}
	return nil
}

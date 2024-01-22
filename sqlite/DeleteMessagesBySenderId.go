package sqlite

import "database/sql"

// deletes messages by SenderId from the MESSAGES table
func DeleteMessagesBySenderId(database *sql.DB, senderId string) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE SenderId = ?", senderId)
	if err != nil {
		return err
	}
	return nil
}

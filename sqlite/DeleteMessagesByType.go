package sqlite

import "database/sql"

// deletes messages of MessageType from the MESSAGES table
func DeleteMessagesByType(database *sql.DB, messageType string) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE MessageType = ?", messageType)
	if err != nil {
		return err
	}
	return nil
}

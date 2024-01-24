package sqlite

import "database/sql"

// Deletes all messages from the MESSAGES table; use with caution
func DeleteAllMessages(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM MESSAGES")
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes message of MessageId from the MESSAGES table
func DeleteMessageById(database *sql.DB, messageId int) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE MessageId = ?", messageId)
	if err != nil {
		return err
	}
	return nil
}

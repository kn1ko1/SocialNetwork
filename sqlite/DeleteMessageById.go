package sqlite

import "database/sql"

// deletes message of MessageId from the MESSAGES table
func DeleteMessagesById(database *sql.DB, messageId string) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE MessageId = ?", messageId)
	if err != nil {
		return err
	}
	return nil
}

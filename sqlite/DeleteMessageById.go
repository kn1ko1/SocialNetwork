package sqlite

import "database/sql"

// deletes message of MessageId from the MESSAGES table
func DeleteMessagesById(db *sql.DB, messageId string) error {
	_, err := db.Exec("DELETE FROM MESSAGES WHERE MessageId = ?", messageId)
	if err != nil {
		return err
	}
	return nil
}

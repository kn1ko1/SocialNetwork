package sqlite

import "database/sql"

// deletes messages by TargetId from the MESSAGES table
func DeleteMessagesByTargetId(db *sql.DB, targetId string) error {
	_, err := db.Exec("DELETE FROM MESSAGES WHERE TargetId = ?", targetId)
	if err != nil {
		return err
	}
	return nil
}

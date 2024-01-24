package sqlite

import "database/sql"

// deletes messages by TargetId from the MESSAGES table
func DeleteMessagesByTargetId(database *sql.DB, targetId int) error {
	_, err := database.Exec("DELETE FROM MESSAGES WHERE TargetId = ?", targetId)
	if err != nil {
		return err
	}
	return nil
}

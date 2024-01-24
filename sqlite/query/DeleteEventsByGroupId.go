package sqlite

import "database/sql"

// deletes events related to GroupId from the EVENT table
func DeleteEventsByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM EVENTS WHERE GroupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

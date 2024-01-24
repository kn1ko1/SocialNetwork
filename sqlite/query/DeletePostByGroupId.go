package sqlite

import "database/sql"

// deletes posts related to groupId from the POSTS table
func DeletePostByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE groupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes posts related to groupId from the POSTS table
func DeletePostByGroupId(db *sql.DB, groupId int) error {
	_, err := db.Exec("DELETE FROM POSTS WHERE groupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes comments related to groupId from the COMMENTS table
func DeleteCommentsByGroupId(db *sql.DB, groupId int) error {
	_, err := db.Exec("DELETE FROM COMMENTS WHERE GroupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

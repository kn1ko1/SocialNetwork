package sqlite

import "database/sql"

// deletes posts related to userId from the POSTS table
func DeletePostByUserId(db *sql.DB, userId int) error {
	_, err := db.Exec("DELETE FROM POSTS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

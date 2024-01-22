package sqlite

import "database/sql"

// deletes posts related to userId from the POSTS table
func DeletePostByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

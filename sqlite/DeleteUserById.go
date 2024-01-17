package sqlite

import "database/sql"

// deletes a specific user from the USERS table
func DeleteUserById(db *sql.DB, userId int) error {
	_, err := db.Exec("DELETE FROM USERS WHERE UserId = ?", userId)
	if err != nil {
		return err
	}
	return nil
}

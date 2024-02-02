package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes post users related to PostId from the POST_USERS table
func DeletePostUsersByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM POST_USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete post users by user ID statement.", err)
		return err
	}
	return nil
}

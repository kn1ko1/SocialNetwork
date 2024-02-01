package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes posts related to userId from the POSTS table
func DeletePostByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete posts by userId statement.", err)
		return err
	}
	return nil
}

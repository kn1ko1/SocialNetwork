package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// deletes a specific user from the USERS table
func DeleteUserById(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete user by ID statement.", err)
		return err
	}
	return nil
}

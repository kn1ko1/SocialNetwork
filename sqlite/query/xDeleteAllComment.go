package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
)

// Deletes all comments from the COMMENTS table; use with caution
func DeleteAllComments(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM COMMENTS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

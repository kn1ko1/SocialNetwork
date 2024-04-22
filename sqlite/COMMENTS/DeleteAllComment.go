package comments

import (
	"database/sql"
	"socialnetwork/utils"
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

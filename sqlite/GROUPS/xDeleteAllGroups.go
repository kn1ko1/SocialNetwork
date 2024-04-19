package groups

import (
	"database/sql"
	"socialnetwork/utils"
)

// Deletes all comments from the COMMENTS table; use with caution
func DeleteAllGroups(database *sql.DB) error {
	_, err := database.Exec("DELETE FROM GROUPS")
	if err != nil {
		utils.HandleError("Error executing delete statement.", err)
		return err
	}
	return nil
}

package groups

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes groupo of GroupId from the GROUPS table
func DeleteGroup(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM GROUPS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing delete all groups statement.", err)
		return err
	}
	return nil
}

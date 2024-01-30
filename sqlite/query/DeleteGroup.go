package sqlite

import "database/sql"

// deletes groupo of GroupId from the GROUPS table
func DeleteGroup(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM GROUPS WHERE GroupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

package sqlite

import "database/sql"

// deletes posts related to groupId from the GROUP_USERS table
func DeleteGroupUserByGroupId(database *sql.DB, groupId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE groupId = ?", groupId)
	if err != nil {
		return err
	}
	return nil
}

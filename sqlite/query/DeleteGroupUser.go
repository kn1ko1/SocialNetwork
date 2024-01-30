package sqlite

import "database/sql"

// deletes groupo of GroupId from the GROUPS table
func DeleteGroupUser(database *sql.DB, groupUserId int) error {
	_, err := database.Exec("DELETE FROM GROUP_USERS WHERE GroupUserId = ?", groupUserId)
	if err != nil {
		return err
	}
	return nil
}

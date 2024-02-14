package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes user users related to subjectId from the USER_USERS table
func DeleteUserUserBySubjectId(database *sql.DB, subjectId int) error {
	_, err := database.Exec("DELETE FROM USER_USERS WHERE SubjectId = ?", subjectId)
	if err != nil {
		utils.HandleError("Error executing delete user users by subject Id.", err)
		return err
	}
	return nil
}

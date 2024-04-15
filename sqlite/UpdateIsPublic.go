package sqlite

import (
	"database/sql"
	"socialnetwork/utils" // Assuming you have a utility package for error handling
)

// UpdateIsPublic updates the IsPublic field for a user in the USERS table
func UpdateIsPublic(database *sql.DB, userId int, isPublic bool) error {
	query := `
		UPDATE USERS
		SET IsPublic = ?
		WHERE UserId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db in UpdateIsPublic.", err)
		return err
	}

	_, err = statement.Exec(isPublic, userId)
	if err != nil {
		utils.HandleError("Error executing statement in UpdateIsPublic.", err)
		return err
	}

	return nil
}

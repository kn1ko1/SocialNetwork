package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Updates user information in the USERS table
func UpdateUser(database *sql.DB, user models.User) (models.User, error) {
	query := `
		UPDATE USERS
		SET
			Bio = ?,
			DOB = ?,
			Email = ?,
			EncryptedPassword = ?,
			FirstName = ?,
			ImageUrl = ?,
			IsPublic = ?,
			LastName = ?,
			UpdatedAt = ?,
			Username = ?
		WHERE UserId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db in UpdateUser.", err)
		return user, err
	}

	_, err = statement.Exec(
		user.Bio,
		user.DOB,
		user.Email,
		user.EncryptedPassword,
		user.FirstName,
		user.ImageUrl,
		user.IsPublic,
		user.LastName,
		user.UpdatedAt,
		user.Username,
		user.UserId,
	)

	if err != nil {
		utils.HandleError("Error executing statement in UpdateUser.", err)

		return user, err
	}

	return user, nil
}

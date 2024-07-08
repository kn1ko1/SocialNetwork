package users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
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
			ImageURL = ?,
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
		user.ImageURL,
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

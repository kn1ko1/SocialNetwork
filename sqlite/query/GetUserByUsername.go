package sqlite

import (
	"database/sql"
	"errors"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Retrieves user with the relevant username from the USERS table
func GetUserByUsername(database *sql.DB, username string) (models.User, error) {
	var user models.User
	err := database.QueryRow("SELECT * FROM USERS WHERE Username = ?", username).
		Scan(
			&user.Bio,
			&user.CreatedAt,
			&user.DOB,
			&user.Email,
			&user.EncryptedPassword,
			&user.FirstName,
			&user.ImageUrl,
			&user.IsPublic,
			&user.LastName,
			&user.UpdatedAt,
			&user.Username,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("user not found.", err)
		return user, errors.New("user not found")
	case err != nil:
		utils.HandleError("Error retrieving user by username.", err)
		return user, err
	}

	return user, nil
}

package users

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves user with the relevant userId from the USERS table
func GetUserByEmail(database *sql.DB, email string) (models.User, error) {
	var user models.User
	err := database.QueryRow("SELECT * FROM USERS WHERE Email = ?", email).
		Scan(
			&user.UserId,
			&user.Bio,
			&user.CreatedAt,
			&user.DOB,
			&user.Email,
			&user.EncryptedPassword,
			&user.FirstName,
			&user.ImageURL,
			&user.IsPublic,
			&user.LastName,
			&user.UpdatedAt,
			&user.Username,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("User not found.", err)
		return user, errors.New("user not found")
	case err != nil:
		utils.HandleError("Error retrieving user by email", err)
		return user, err
	}

	return user, nil
}

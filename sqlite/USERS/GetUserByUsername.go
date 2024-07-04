package users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
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
			&user.ImageURL,
			&user.IsPublic,
			&user.LastName,
			&user.UpdatedAt,
			&user.Username,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("user not found.", sql.ErrNoRows)
		return user, sql.ErrNoRows
	case err != nil:
		utils.HandleError("Error retrieving user by username.", err)
		return user, err
	}

	return user, nil
}

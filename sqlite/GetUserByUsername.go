package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves user with the relevant username from the USERS table
func GetUserByUsername(database *sql.DB, username string) (*models.User, error) {
	var user models.User
	err := database.QueryRow("SELECT * FROM USERS WHERE Username = ?", username).
		Scan(
			&user.Bio,
			&user.CreatedAt,
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
		return nil, errors.New("user not found")
	case err != nil:
		return nil, err
	}

	return &user, nil
}

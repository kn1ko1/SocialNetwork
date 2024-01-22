package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves user with the relevant userId from the USERS table
func GetUserByEmail(database *sql.DB, email string) (models.User, error) {
	var user models.User
	err := database.QueryRow("SELECT * FROM USERS WHERE Email = ?", email).
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
		return user, errors.New("user not found")
	case err != nil:
		return user, err
	}

	return user, nil
}

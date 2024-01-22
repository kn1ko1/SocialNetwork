package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves user with the relevant userId from the USERS table
func GetUserById(database *sql.DB, userId int) (models.User, error) {
	var user models.User
	err := database.QueryRow("SELECT * FROM USERS WHERE UserId = ?", userId).
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

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves all users from the USERS table
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.UserId,
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
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

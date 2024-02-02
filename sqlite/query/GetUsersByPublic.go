package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves public users from the USERS table
func GetUsersByPublic(database *sql.DB) ([]models.User, error) {
	var users []models.User

	rows, err := database.Query("SELECT * FROM USERS WHERE IsPublic = true")
	if err != nil {
		utils.HandleError("Error querying public users", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(
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
			utils.HandleError("Error scanning row in GetPublicUsers", err)
			return users, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPublicUsers", err)
		return users, err
	}

	return users, nil
}

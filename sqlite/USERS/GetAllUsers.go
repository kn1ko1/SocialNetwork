package users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Retrieves all users from the USERS table
func GetAllUsers(database *sql.DB) ([]models.User, error) {
	rows, err := database.Query("SELECT * FROM USERS")
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM USERS statement.", err)
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
			&user.ImageURL,
			&user.IsPublic,
			&user.LastName,
			&user.UpdatedAt,
			&user.Username,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllUsers.", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllUsers.", err)
		return nil, err
	}

	return users, nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves postUsers with the relevant userId from the POST_USERS table
func GetPostUsersByUserId(database *sql.DB, userId int) ([]models.PostUser, error) {
	rows, err := database.Query("SELECT * FROM POST_USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error querying postUsers by UserId.", err)
		return nil, err
	}
	defer rows.Close()

	var postUsers []models.PostUser

	for rows.Next() {
		var postUser models.PostUser
		err := rows.Scan(
			&postUser.PostUserId,
			&postUser.CreatedAt,
			&postUser.PostId,
			&postUser.UpdatedAt,
			&postUser.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning row in GetPostUsersByUserId.", err)
			return nil, err
		}

		postUsers = append(postUsers, postUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostUsersByUserId.", err)
		return nil, err
	}

	return postUsers, nil
}

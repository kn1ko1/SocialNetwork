package post_users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves postUsers with the relevant postId from the POST_USERS table
func GetPostUsersByPostId(database *sql.DB, postId int) ([]models.PostUser, error) {
	rows, err := database.Query("SELECT * FROM POST_USERS WHERE PostId = ?", postId)
	if err != nil {
		utils.HandleError("Error querying postUsers by PostId.", err)
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
			utils.HandleError("Error scanning row in GetPostUsersByPostId.", err)
			return nil, err
		}

		postUsers = append(postUsers, postUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostUsersByPostId.", err)
		return nil, err
	}

	return postUsers, nil
}

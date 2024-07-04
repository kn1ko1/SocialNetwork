package user_users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Retrieves userUsers with the relevant subjectId from the USER_USERS table
func GetUserUsersByFollowerId(database *sql.DB, followerId int) ([]models.UserUser, error) {
	rows, err := database.Query("SELECT * FROM USER_USERS WHERE FollowerId = ?", followerId)
	if err != nil {
		utils.HandleError("Error querying userUsers by FollowerId.", err)
		return nil, err
	}
	defer rows.Close()

	var userUsers []models.UserUser

	for rows.Next() {
		var userUser models.UserUser
		err := rows.Scan(
			&userUser.UserUserId,
			&userUser.CreatedAt,
			&userUser.FollowerId,
			&userUser.SubjectId,
			&userUser.UpdatedAt,
		)
		if err != nil {
			utils.HandleError("Error scanning row in GetUserUsersByFollowerId.", err)
			return nil, err
		}

		userUsers = append(userUsers, userUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetUserUsersByFollowerId.", err)
		return nil, err
	}

	return userUsers, nil
}

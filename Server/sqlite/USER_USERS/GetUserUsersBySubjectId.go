package user_users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves userUsers with the relevant subjectId from the USER_USERS table
func GetUserUsersBySubjectId(database *sql.DB, subjectId int) ([]models.UserUser, error) {
	rows, err := database.Query("SELECT * FROM USER_USERS WHERE SubjectId = ?", subjectId)
	if err != nil {
		utils.HandleError("Error querying userUsers by SubjectId.", err)
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
			utils.HandleError("Error scanning row in GetUserUsersBySubjectId.", err)
			return nil, err
		}

		userUsers = append(userUsers, userUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetUserUsersBySubjectId.", err)
		return nil, err
	}

	return userUsers, nil
}

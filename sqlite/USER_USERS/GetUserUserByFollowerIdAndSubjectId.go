package user_users

import (
	"database/sql"
	"errors"
	"log"
	"socialnetwork/models"
	"socialnetwork/utils"
)

func GetUserUserByFollowerIdAndSubjectId(database *sql.DB, followerId, subjectId int) (models.UserUser, error) {
	var userUser models.UserUser
	queryStr := `SELECT * FROM USER_USERS WHERE FollowerId = ? AND SubjectId = ?;`

	// Execute the SELECT query
	row := database.QueryRow(queryStr, followerId, subjectId)

	// Scan the result into variables
	err := row.Scan(
		&userUser.UserUserId,
		&userUser.CreatedAt,
		&userUser.FollowerId,
		&userUser.SubjectId,
		&userUser.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("UserUser not found", err)
		return userUser, errors.New("UserUser not found")
	case err != nil:
		utils.HandleError("Error retrieving UserUser", err)
		return userUser, err
	}
	log.Println("found UserUser in GetUserUserByFollowerIdAndSubjectId")
	return userUser, nil
}

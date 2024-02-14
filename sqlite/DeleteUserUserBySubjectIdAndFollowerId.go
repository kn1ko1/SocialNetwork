package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

func DeleteUserUserBySubjectIdAndFollowerId(database *sql.DB, subjectId, followerId int) error {

	queryStr := `DELETE *
	FROM USER_USERS
	WHERE (SubjectId = (?) AND FollowerId = (?))
	OR (FollowerId = (?) AND SubjectId = (?))
	ORDER BY timestamp ASC;`

	_, err := database.Query(queryStr, subjectId, followerId, followerId, subjectId)
	if err != nil {
		utils.HandleError("Error executing delete user users by subject Id and follower ID statement.", err)
		return err
	}
	return nil
}

package user_users

import (
	"database/sql"
	"socialnetwork/utils"
)

func DeleteUserUserBySubjectIdAndFollowerId(database *sql.DB, subjectId, followerId int) error {
	queryStr := `DELETE FROM USER_USERS
        WHERE SubjectId = ? AND FollowerId = ?;`

	_, err := database.Exec(queryStr, subjectId, followerId)
	if err != nil {
		utils.HandleError("Error executing delete user users by subject Id and follower ID statement.", err)
		return err
	}
	return nil
}

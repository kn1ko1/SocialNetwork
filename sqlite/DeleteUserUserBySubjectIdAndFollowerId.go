package sqlite

import (
	"database/sql"
	"log"
	"socialnetwork/utils"
)

func DeleteUserUserBySubjectIdAndFollowerId(database *sql.DB, subjectId, followerId int) error {
	log.Println("subjectId, followerId", subjectId, followerId)
	queryStr := `DELETE FROM USER_USERS
        WHERE SubjectId = ? AND FollowerId = ?;`

	_, err := database.Exec(queryStr, subjectId, followerId)
	if err != nil {
		utils.HandleError("Error executing delete user users by subject Id and follower ID statement.", err)
		return err
	}
	return nil
}

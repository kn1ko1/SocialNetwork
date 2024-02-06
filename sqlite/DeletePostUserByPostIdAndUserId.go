package sqlite

import (
	"database/sql"
	"socialnetwork/utils"
)

func DeletePostUserByPostIdAndUserId(database *sql.DB, userId, postId int) error {

	queryStr := `DELETE FROM POST_USERS
	WHERE (UserId = ? AND PostId = ?)
	OR (PostId = ? AND UserId = ?);`

	_, err := database.Exec(queryStr, userId, postId, postId, userId)
	if err != nil {
		utils.HandleError("Error executing delete post user statement.", err)
		return err
	}
	return nil
}

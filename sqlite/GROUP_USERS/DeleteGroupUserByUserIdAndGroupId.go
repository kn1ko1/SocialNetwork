package group_users

import (
	"database/sql"
	"socialnetwork/utils"
)

func DeleteGroupUserByGroupIdAndUserId(database *sql.DB, groupId, userId int) error {

	queryStr := `DELETE *
	FROM GROUP_USERS
	WHERE (UserId = (?) AND GroupId = (?))
	OR (GroupId = (?) AND UserId = (?))
	ORDER BY timestamp ASC;`

	_, err := database.Query(queryStr, userId, groupId, groupId, userId)
	if err != nil {
		utils.HandleError("Error executing delete group users by group Id and user ID statement.", err)
		return err
	}
	return nil
}

package sqlite

import (
	"database/sql"
)

func DeleteGroupUserByUserIdAndGroupId(database *sql.DB, userId, groupId int) error {

	queryStr := `DELETE *
	FROM GROUP_USERS
	WHERE (UserId = (?) AND GroupId = (?))
	OR (GroupId = (?) AND UserId = (?))
	ORDER BY timestamp ASC;`

	_, err := database.Query(queryStr, userId, groupId, groupId, userId)
	if err != nil {
		return err
	}
	return nil
}

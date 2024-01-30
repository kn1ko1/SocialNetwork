package sqlite

import (
	"database/sql"
)

func DeleteGroupUserByUserIdAndGroupId(database *sql.DB, UserId, GroupId int) error {

	queryStr := `DELETE *
	FROM GROUPUSERS
	WHERE (UserId = (?) AND GroupId = (?))
	OR (GroupId = (?) AND UserId = (?))
	ORDER BY timestamp ASC;`

	_, err := database.Query(queryStr, UserId, GroupId, GroupId, UserId)
	if err != nil {
		return err
	}
	return nil
}

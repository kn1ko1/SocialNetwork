package sqlite

import (
	"database/sql"
)

func DeleteEventUserByUserIdAndEventId(database *sql.DB, userId, eventId int) error {

	queryStr := `DELETE *
	FROM EVENT_USERS
	WHERE (UserId = (?) AND GroupId = (?))
	OR (GroupId = (?) AND UserId = (?))
	ORDER BY timestamp ASC;`

	_, err := database.Query(queryStr, userId, eventId, eventId, userId)
	if err != nil {
		return err
	}
	return nil
}

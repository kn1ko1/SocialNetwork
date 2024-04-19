package event_users

import (
	"database/sql"
	"socialnetwork/utils"
)

func DeleteEventUserByEventIdAndUserId(database *sql.DB, userId, eventId int) error {

	queryStr := `DELETE FROM EVENT_USERS
	WHERE (UserId = ? AND EventId = ?)
	OR (EventId = ? AND UserId = ?);`

	_, err := database.Exec(queryStr, userId, eventId, eventId, userId)
	if err != nil {
		utils.HandleError("Error executing delete event user statement.", err)
		return err
	}
	return nil
}

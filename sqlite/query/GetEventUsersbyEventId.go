package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

func GetEventUsersByEventId(database *sql.DB, eventId int) ([]models.EventUser, error) {
	rows, err := database.Query("SELECT * FROM EVENT_USERS WHERE EventId = ?", eventId)
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM EVENT_USERS WHERE EventId = ? statement.", err)
		return nil, err
	}
	defer rows.Close()

	var eventUsers []models.EventUser

	for rows.Next() {
		var eventUser models.EventUser
		err := rows.Scan(
			&eventUser.EventUserId,
			&eventUser.CreatedAt,
			&eventUser.UpdatedAt,
			&eventUser.UserId,
			&eventUser.EventId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetEventUsersByEventId.", err)
			return nil, err
		}

		eventUsers = append(eventUsers, eventUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetEventUsersByEventId.", err)
		return nil, err
	}

	return eventUsers, nil
}

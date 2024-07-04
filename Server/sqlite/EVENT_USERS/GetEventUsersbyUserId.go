package event_users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves event with the relevant userId from the EVENTS table
func GetEventUsersByUserId(database *sql.DB, userId int) ([]models.EventUser, error) {
	rows, err := database.Query("SELECT * FROM EVENT_USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM EVENT_USERS WHERE UserId = ? statement.", err)
		return nil, err
	}
	defer rows.Close()

	var eventUsers []models.EventUser

	for rows.Next() {
		var eventUser models.EventUser
		err := rows.Scan(
			&eventUser.EventUserId,
			&eventUser.CreatedAt,
			&eventUser.EventId,
			&eventUser.IsGoing,
			&eventUser.UpdatedAt,
			&eventUser.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetEventUsersByUserId.", err)
			return nil, err
		}

		eventUsers = append(eventUsers, eventUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetEventUsersByUserId.", err)
		return nil, err
	}

	return eventUsers, nil
}

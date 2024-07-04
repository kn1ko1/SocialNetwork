package events

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves event with the relevant userId from the EVENTS table
func GetEventsByUserId(database *sql.DB, userId int) ([]models.Event, error) {
	var events []models.Event
	rows, err := database.Query("SELECT * FROM EVENTS WHERE UserId = ?", userId)
	if err != nil {
		// no results in DB
		return events, nil
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.EventId,
			&event.CreatedAt,
			&event.DateTime,
			&event.Description,
			&event.GroupId,
			&event.Title,
			&event.UpdatedAt,
			&event.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetEventsByUserId.", err)
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetEventsByUserId.", err)
		return nil, err
	}

	return events, nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves all events from the EVENTS table
func GetAllEvents(database *sql.DB) ([]models.Event, error) {
	rows, err := database.Query("SELECT * FROM EVENTS")
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM EVENTS statement.", err)
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.EventId,
			&event.CreatedAt,
			&event.DateTime,
			&event.Description,
			&event.GroupId,
			&event.UpdatedAt,
			&event.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllEvents.", err)
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllEvents.", err)
		return nil, err
	}

	return events, nil
}

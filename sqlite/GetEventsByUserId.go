package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves event with the relevant userId from the EVENTS table
func GetEventsByUserId(database *sql.DB, userId int) ([]*models.Event, error) {
	rows, err := database.Query("SELECT * FROM EVENTS WHERE UserId = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*models.Event

	for rows.Next() {
		var event *models.Event
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
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

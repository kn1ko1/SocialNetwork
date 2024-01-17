package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves event with the relevant groupId from the EVENTS table
func GetEventsByGroupId(database *sql.DB, groupId int) ([]*models.Event, error) {
	rows, err := database.Query("SELECT * FROM EVENTS WHERE GroupId = ?", groupId)
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

package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves event with the relevant eventId from the EVENTS table
func GetEventById(database *sql.DB, eventId int) (models.Event, error) {
	var event models.Event
	err := database.QueryRow("SELECT * FROM EVENTS WHERE EventId = ?", eventId).
		Scan(
			&event.EventId,
			&event.CreatedAt,
			&event.DateTime,
			&event.Description,
			&event.GroupId,
			&event.Title,
			&event.UpdatedAt,
			&event.UserId,
		)

	switch {
	case err == sql.ErrNoRows:
		return event, errors.New("event not found")
	case err != nil:
		return event, err
	}

	return event, nil
}

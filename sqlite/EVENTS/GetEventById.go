package events

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
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
		utils.HandleError("Event not found in GetEventById.", err)
		return event, errors.New("event not found")
	case err != nil:
		utils.HandleError("Error retrieving event by ID in GetEventById.", err)
		return event, err
	}

	return event, nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Updates event information in the EVENTS table
func UpdateEvent(database *sql.DB, event models.Event) (models.Event, error) {
	query := `
		UPDATE EVENTS
		SET
			DateTime = ?,
			Description = ?,
			Title = ?,
			UpdatedAt = ?
		WHERE EventId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		return event, err
	}

	_, err = statement.Exec(
		event.DateTime,
		event.Description,
		event.Title,
		event.UpdatedAt,
		event.EventId,
	)

	if err != nil {
		return event, err
	}

	return event, nil
}

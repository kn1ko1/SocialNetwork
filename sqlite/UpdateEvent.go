package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils" // Assuming you have a utility package for error handling
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
		utils.HandleError("Error preparing statement in UpdateEvent.", err)
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
		utils.HandleError("Error executing query in UpdateEvent.", err)
		return event, err
	}

	return event, nil
}

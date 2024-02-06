package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Adds event into the given database
func CreateEvent(database *sql.DB, event models.Event) (models.Event, error) {

	query := `
	INSERT INTO EVENTS (
		CreatedAt,
		DateTime,
		Description,
		GroupId,
		Title,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return event, err
	}
	res, err := statement.Exec(
		event.CreatedAt,
		event.DateTime,
		event.Description,
		event.GroupId,
		event.Title,
		event.UpdatedAt,
		event.UserId,
	)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return event, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return event, err
	}
	event.EventId = int(id)
	return event, nil
}

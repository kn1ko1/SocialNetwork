package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds eventUser into the given database
func CreateEventUser(database *sql.DB, eventUser models.EventUser) (models.EventUser, error) {

	query := `
	INSERT INTO EVENT_USERS (
		CreatedAt,
		EventId,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		return eventUser, err
	}
	res, err := statement.Exec(
		eventUser.CreatedAt,
		eventUser.EventId,
		eventUser.UpdatedAt,
		eventUser.UserId,
	)
	if err != nil {
		return eventUser, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return eventUser, err
	}
	eventUser.EventUserId = int(id)
	return eventUser, nil
}

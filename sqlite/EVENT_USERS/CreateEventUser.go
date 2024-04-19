package event_users

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Adds eventUser into the given database
func CreateEventUser(database *sql.DB, eventUser models.EventUser) (models.EventUser, error) {

	query := `
	INSERT INTO EVENT_USERS (
		CreatedAt,
		EventId,
		IsGoing,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return eventUser, err
	}
	res, err := statement.Exec(
		eventUser.CreatedAt,
		eventUser.EventId,
		eventUser.IsGoing,
		eventUser.UpdatedAt,
		eventUser.UserId,
	)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return eventUser, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return eventUser, err
	}
	eventUser.EventUserId = int(id)
	return eventUser, nil
}

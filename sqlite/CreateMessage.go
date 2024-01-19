package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds message into the given database
func CreateMessage(database *sql.DB, message models.Message) (models.Message, error) {

	query := `
	INSERT INTO MESSAGES (
		Body,
		CreatedAt,
		MessageType,
		SenderId,
		TargetId,
		UpdatedAt
	) VALUES (?, ?, ?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		return message, err
	}
	res, err := statement.Exec(
		message.Body,
		message.CreatedAt,
		message.MessageType,
		message.SenderId,
		message.TargetId,
		message.UpdatedAt,
	)
	if err != nil {
		return message, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return message, err
	}
	message.MessageId = int(id)
	return message, nil
}

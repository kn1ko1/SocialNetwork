package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	// Assuming you have a utility package for error handling
)

// Updates message information in the MESSAGES table
func UpdateMessage(database *sql.DB, message models.Message) (models.Message, error) {
	query := `
		UPDATE MESSAGES
		SET
			Body = ?,
			MessageType = ?,
			SenderId = ?,
			TargetId = ?,
			UpdatedAt = ?
		WHERE MessageId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		return message, err
	}

	_, err = statement.Exec(
		message.Body,
		message.MessageType,
		message.SenderId,
		message.TargetId,
		message.UpdatedAt,
		message.MessageId,
	)

	if err != nil {
		return message, err
	}

	return message, nil
}

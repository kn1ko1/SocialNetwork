// Package messages does something.
package messages

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
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
		utils.HandleError("Error preparing db query.", err)
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
		utils.HandleError("Error executing statement.", err)
		return message, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return message, err
	}
	message.MessageId = int(id)
	return message, nil
}

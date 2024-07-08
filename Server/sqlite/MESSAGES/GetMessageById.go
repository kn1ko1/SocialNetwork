package messages

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves message with the relevant messageId from the MESSAGES table
func GetMessageById(database *sql.DB, messageId int) (models.Message, error) {
	var message models.Message
	err := database.QueryRow("SELECT * FROM MESSAGES WHERE MessageId = ?", messageId).
		Scan(
			&message.MessageId,
			&message.Body,
			&message.CreatedAt,
			&message.MessageType,
			&message.SenderId,
			&message.TargetId,
			&message.UpdatedAt,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("Message not found in GetMessageById.", err)
		return message, errors.New("message not found")
	case err != nil:
		utils.HandleError("Error retrieving message by ID in GetMessageById.", err)
		return message, err
	}

	return message, nil
}

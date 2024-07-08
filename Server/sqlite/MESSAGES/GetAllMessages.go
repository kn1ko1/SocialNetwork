package messages

import (
	"database/sql"

	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves all messages from the MESSAGES table
func GetAllMessages(database *sql.DB) ([]models.Message, error) {
	rows, err := database.Query("SELECT * FROM MESSAGES")
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM MESSAGES statement.", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var message models.Message
		err := rows.Scan(
			&message.MessageId,
			&message.Body,
			&message.CreatedAt,
			&message.MessageType,
			&message.SenderId,
			&message.TargetId,
			&message.UpdatedAt,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllMessages.", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllMessages.", err)
		return nil, err
	}

	return messages, nil
}

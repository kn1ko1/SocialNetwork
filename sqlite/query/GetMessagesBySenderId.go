package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves messages with the relevant senderId from the MESSAGES table
func GetMessagesBySenderId(database *sql.DB, senderId int) ([]models.Message, error) {
	rows, err := database.Query("SELECT * FROM MESSAGES WHERE SenderId = ?", senderId)
	if err != nil {
		utils.HandleError("Error executing query in GetMessagesBySenderId.", err)
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
			utils.HandleError("Error scanning rows in GetMessagesBySenderId.", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetMessagesBySenderId.", err)
		return nil, err
	}

	return messages, nil
}

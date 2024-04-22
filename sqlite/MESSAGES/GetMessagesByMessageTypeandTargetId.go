package messages

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves messages with the relevant messageType and TargetId from the MESSAGES table
func GetMessagesByMessageTypeandTargetId(database *sql.DB, messageType string, targetId int) ([]models.Message, error) {

	rows, err := database.Query("SELECT * FROM MESSAGES WHERE MessageType = ? AND TargetId = ?", messageType, targetId)
	if err != nil {
		utils.HandleError("Error executing query in GetMessagesByType.", err)
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
			utils.HandleError("Error scanning rows in GetMessagesByType.", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetMessagesByType.", err)
		return nil, err
	}

	return messages, nil
}

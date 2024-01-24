package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves message with the relevant senderId from the MESSAGES table
func GetMessagesBySenderId(database *sql.DB, senderId int) ([]models.Message, error) {
	rows, err := database.Query("SELECT * FROM MESSAGES WHERE SenderId = ?", senderId)
	if err != nil {
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
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

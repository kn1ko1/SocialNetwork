package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves message with the relevant messageType from the MESSAGES table
func GetMessagesByType(database *sql.DB, messageType string) ([]*models.Message, error) {
	rows, err := database.Query("SELECT * FROM MESSAGES WHERE MessageType = ?", messageType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message

	for rows.Next() {
		var message *models.Message
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

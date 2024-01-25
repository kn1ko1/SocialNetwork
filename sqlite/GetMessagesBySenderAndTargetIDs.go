package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves message with the relevant senderId from the MESSAGES table
func GetMessagesBySenderAndTargetIDs(database *sql.DB, senderId, targetId int) ([]models.Message, error) {

	queryStr := `SELECT *
	FROM "messages"
	WHERE (SenderID = (?) AND TargetID = (?))
	OR (SenderID = (?) AND TargetID = (?))
	ORDER BY timestamp ASC;`

	var messages []models.Message
	rows, err := database.Query(queryStr, senderId, targetId, targetId, senderId)
	if err != nil {
		return messages, err
	}
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

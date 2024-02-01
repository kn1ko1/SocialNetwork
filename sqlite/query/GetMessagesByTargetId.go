package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves messages with the relevant targetId from the MESSAGES table
func GetMessagesByTargetId(database *sql.DB, targetId int) ([]models.Message, error) {
	rows, err := database.Query("SELECT * FROM MESSAGES WHERE TargetId = ?", targetId)
	if err != nil {
		utils.HandleError("Error executing query in GetMessagesByTargetId.", err)
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
			utils.HandleError("Error scanning rows in GetMessagesByTargetId.", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetMessagesByTargetId.", err)
		return nil, err
	}

	return messages, nil
}

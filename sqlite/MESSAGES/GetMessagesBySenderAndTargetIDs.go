package messages

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Retrieves messages with the relevant senderId and targetId from the MESSAGES table
func GetMessagesBySenderAndTargetIds(database *sql.DB, senderId, targetId int) ([]models.Message, error) {

	queryStr := `SELECT *
	FROM MESSAGES
	WHERE (SenderId = (?) AND TargetId = (?))
	OR (SenderId = (?) AND TargetId = (?))
	ORDER BY CreatedAt Desc;`

	var messages []models.Message
	rows, err := database.Query(queryStr, senderId, targetId, targetId, senderId)
	if err != nil {
		utils.HandleError("Error executing query in GetMessagesBySenderAndTargetIds.", err)
		return messages, err
	}
	defer rows.Close()

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
			utils.HandleError("Error scanning rows in GetMessagesBySenderAndTargetIds.", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetMessagesBySenderAndTargetIds.", err)
		return nil, err
	}

	return messages, nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves all events from the EVENTS table
func GetAllGroups(database *sql.DB) ([]models.Group, error) {
	rows, err := database.Query("SELECT * FROM GROUPS")
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM GROUPS statement.", err)
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group

	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.GroupId,
			&group.CreatedAt,
			&group.Description,
			&group.Title,
			&group.UpdatedAt,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllGroups.", err)
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllGroups.", err)
		return nil, err
	}

	return groups, nil
}

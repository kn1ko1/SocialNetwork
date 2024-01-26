package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves all events from the EVENTS table
func GetAllGroups(database *sql.DB) ([]models.Group, error) {
	rows, err := database.Query("SELECT * FROM GROUPS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group

	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.GroupID,
			&group.CreatedAt,
			&group.Description,
			&group.Title,
			&group.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

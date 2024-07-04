package groups

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// GetGroupsByUserId retrieves groups associated with a given user ID using foreign keys
func GetGroupsByUserId(database *sql.DB, userId int) ([]models.Group, error) {
	var groups []models.Group

	// Query to select groups based on the provided userId
	query := `
        SELECT g.GroupId, g.CreatedAt, g.CreatorId, g.Description, g.Title, g.UpdatedAt
        FROM GROUP_USERS gu
        JOIN GROUPS g ON gu.GroupId = g.GroupId
        WHERE gu.UserId = ?
    `

	// Execute the query
	rows, err := database.Query(query, userId)
	if err != nil {
		// no results in DB
		return groups, nil
	}
	defer rows.Close()

	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.GroupId,
			&group.CreatedAt,
			&group.CreatorId,
			&group.Description,
			&group.Title,
			&group.UpdatedAt,
		)
		if err != nil {
			utils.HandleError("Error scanning row in GetGroupsByUserId.", err)
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetGroupsByUserId.", err)
		return nil, err
	}

	return groups, nil
}

// func GetGroupsByUserId(database *sql.DB, userId int) ([]models.Group, error) {
// 	var groups []models.Group
// 	groupUsers, err := GetGroupUsersByUserId(database, userId)
// 	if err != nil {
// 		utils.HandleError("Problem.", err)
// 	}

// 	for i := 0; i < len(groupUsers)-1; i++ {
// 		group, _ := GetGroupById(database, groupUsers[i].GroupId)
// 		groups = append(groups, group)

// 	}
// 	return groups, nil
// }

package groups

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves group with the relevant groupId from the GROUPS table
func GetGroupById(database *sql.DB, groupId int) (models.Group, error) {
	var group models.Group
	err := database.QueryRow("SELECT * FROM GROUPS WHERE GroupId = ?", groupId).
		Scan(
			&group.GroupId,
			&group.CreatedAt,
			&group.CreatorId,
			&group.Description,
			&group.Title,
			&group.UpdatedAt,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("Group not found in GetGroup.", err)
		return group, errors.New("group not found")
	case err != nil:
		utils.HandleError("Error retrieving group by ID in GetGroup.", err)
		return group, err
	}

	return group, nil
}

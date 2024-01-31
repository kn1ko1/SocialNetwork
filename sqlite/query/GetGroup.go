package sqlite

import (
	"database/sql"
	"errors"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Retrieves group with the relevant groupId from the GROUPS table
func GetGroup(database *sql.DB, groupId int) (models.Group, error) {
	var group models.Group
	err := database.QueryRow("SELECT * FROM GROUPS WHERE GroupId = ?", groupId).
		Scan(
			&group.CreatedAt,
			&group.CreatorId,
			&group.Description,
			&group.Title,
			&group.UpdatedAt,
			&group.GroupId,
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

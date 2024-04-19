package group_users

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves groupUser with the relevant groupUserId from the GROUP_USERS table
func GetGroupUser(database *sql.DB, groupUserId int) (models.GroupUser, error) {
	var groupUser models.GroupUser
	err := database.QueryRow("SELECT * FROM GROUP_USERS WHERE GroupUserId = ?", groupUserId).
		Scan(
			&groupUser.CreatedAt,
			&groupUser.GroupId,
			&groupUser.UpdatedAt,
			&groupUser.UserId,
			&groupUser.GroupUserId,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("GroupUser not found in GetGroupUser.", err)
		return groupUser, errors.New("groupUser not found")
	case err != nil:
		utils.HandleError("Error retrieving groupUser by ID in GetGroupUser.", err)
		return groupUser, err
	}

	return groupUser, nil
}

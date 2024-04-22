package group_users

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

func GetGroupUsersByGroupId(database *sql.DB, groupId int) ([]models.GroupUser, error) {
	rows, err := database.Query("SELECT * FROM GROUP_USERS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM GROUP_USERS WHERE GroupId = ? statement.", err)
		return nil, err
	}
	defer rows.Close()

	var groupUsers []models.GroupUser

	for rows.Next() {
		var groupUser models.GroupUser
		err := rows.Scan(
			&groupUser.GroupUserId,
			&groupUser.CreatedAt,
			&groupUser.GroupId,
			&groupUser.UpdatedAt,
			&groupUser.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetGroupUsersByGroupId.", err)
			return nil, err
		}

		groupUsers = append(groupUsers, groupUser)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetGroupUsersByGroupId.", err)
		return nil, err
	}

	return groupUsers, nil
}

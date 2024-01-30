package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

func GetGroupUsersByGroupId(database *sql.DB, groupId int) ([]models.GroupUser, error) {
	rows, err := database.Query("SELECT * FROM GROUP_USERS WHERE GroupId = ?", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groupUsers []models.GroupUser

	for rows.Next() {
		var groupUser models.GroupUser
		err := rows.Scan(
			&groupUser.GroupUserId,
			&groupUser.CreatedAt,
			&groupUser.UpdatedAt,
			&groupUser.UserId,
			&groupUser.GroupId,
		)
		if err != nil {
			return nil, err
		}

		groupUsers = append(groupUsers, groupUser)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groupUsers, nil
}

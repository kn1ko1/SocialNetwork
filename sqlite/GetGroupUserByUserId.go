package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves event with the relevant userId from the EVENTS table
func GetGroupUsersByUserId(database *sql.DB, userId int) ([]models.GroupUser, error) {
	rows, err := database.Query("SELECT * FROM GROUP_USERS WHERE UserId = ?", userId)
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
			&groupUser.GroupId,
			&groupUser.UpdatedAt,
			&groupUser.UserId,
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

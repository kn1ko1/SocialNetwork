package group_users

import (
	"database/sql"
	"log"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves event with the relevant userId from the EVENTS table
func GetGroupUsersByUserId(database *sql.DB, userId int) ([]models.GroupUser, error) {
	rows, err := database.Query("SELECT * FROM GROUP_USERS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM GROUP_USERS WHERE UserId = ? statement.", err)
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
			utils.HandleError("Error scanning rows in GetGroupUsersByUserId.", err)
			return nil, err
		}

		groupUsers = append(groupUsers, groupUser)
	}
	log.Println("[GetGroupUsersByUserId], groupUsers:", groupUsers)
	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetGroupUsersByUserId.", err)
		return nil, err
	}

	return groupUsers, nil
}

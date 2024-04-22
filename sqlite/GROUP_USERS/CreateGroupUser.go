package group_users

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Adds groupUser into the given database
func CreateGroupUser(database *sql.DB, groupUser models.GroupUser) (models.GroupUser, error) {

	query := `
	INSERT INTO GROUP_USERS (
		CreatedAt,
		GroupId,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return groupUser, err
	}
	res, err := statement.Exec(
		groupUser.CreatedAt,
		groupUser.GroupId,
		groupUser.UpdatedAt,
		groupUser.UserId,
	)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return groupUser, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)

		return groupUser, err
	}
	groupUser.GroupUserId = int(id)
	return groupUser, nil
}

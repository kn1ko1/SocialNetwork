package user_users

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Adds post user into the given database
func CreateUserUser(database *sql.DB, userUser models.UserUser) (models.UserUser, error) {

	query := `
	INSERT INTO USER_USERS (
		CreatedAt,
		FollowerId,
		SubjectId,
		UpdatedAt	
	) VALUES (?, ?, ?, ?)
`
	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return userUser, err
	}

	res, err := statement.Exec(
		userUser.CreatedAt,
		userUser.FollowerId,
		userUser.SubjectId,
		userUser.UpdatedAt)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return userUser, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return userUser, err
	}

	userUser.UserUserId = int(id)
	return userUser, nil
}

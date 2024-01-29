package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds groupUser into the given database
func CreateGroupUser(database *sql.DB, groupUser models.GroupUser) (models.GroupUser, error) {

	query := `
	INSERT INTO GROUPS (
		CreatedAt,
		GroupId,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		return groupUser, err
	}
	res, err := statement.Exec(
		groupUser.CreatedAt,
		groupUser.GroupId,
		groupUser.UpdatedAt,
		groupUser.UpdatedAt,
		groupUser.UserId,
	)
	if err != nil {
		return groupUser, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return groupUser, err
	}
	groupUser.GroupUserId = int(id)
	return groupUser, nil
}

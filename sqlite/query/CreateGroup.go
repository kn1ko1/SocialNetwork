package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds group into the given database
func CreateGroup(database *sql.DB, group models.Group) (models.Group, error) {

	query := `
	INSERT INTO GROUPS (
		CreatedAt,
		CreatorId,
		Description,
		Title,
		UpdatedAt
	) VALUES (?, ?, ?, ?, ?)
`

	statement, err := database.Prepare(query)
	if err != nil {
		return group, err
	}
	res, err := statement.Exec(
		group.CreatedAt,
		group.CreatorId,
		group.Description,
		group.Title,
		group.UpdatedAt,
	)
	if err != nil {
		return group, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return group, err
	}
	group.GroupId = int(id)
	return group, nil
}

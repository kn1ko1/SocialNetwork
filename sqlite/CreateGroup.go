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
		CreatorID,
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
		group.CreatorID,
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
	group.GroupID = int(id)
	return group, nil
}

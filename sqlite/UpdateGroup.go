package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Updates group information in the GROUPS table
func UpdateGroup(database *sql.DB, group models.Group) (models.Group, error) {
	query := `
		UPDATE GROUPS
		SET
			Description = ?,
			Title = ?,
			UpdatedAt = ?
		WHERE GroupID = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		return group, err
	}

	_, err = statement.Exec(
		group.Description,
		group.Title,
		group.UpdatedAt,
		group.GroupId,
	)

	if err != nil {
		return group, err
	}

	return group, nil
}

package groups

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils" // Assuming you have a utility package for error handling
)

// Updates group information in the GROUPS table
func UpdateGroup(database *sql.DB, group models.Group) (models.Group, error) {
	query := `
		UPDATE GROUPS
		SET
			Description = ?,
			Title = ?,
			UpdatedAt = ?
		WHERE GroupId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing statement in UpdateGroup.", err)
		return group, err
	}

	_, err = statement.Exec(
		group.Description,
		group.Title,
		group.UpdatedAt,
		group.GroupId,
	)

	if err != nil {
		utils.HandleError("Error executing query in UpdateGroup.", err)
		return group, err
	}

	return group, nil
}

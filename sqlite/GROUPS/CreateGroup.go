package groups

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
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
		utils.HandleError("Error preparing db query.", err)

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
		utils.HandleError("Error executing statement.", err)
		return group, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return group, err
	}
	group.GroupId = int(id)
	return group, nil
}

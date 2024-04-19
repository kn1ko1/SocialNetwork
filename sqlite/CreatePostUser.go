package sqlite

import (
	"database/sql"
	"log"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Adds post user into the given database
func CreatePostUser(database *sql.DB, postUser models.PostUser) (models.PostUser, error) {

	query := `
	INSERT INTO POST_USERS (
		CreatedAt,
		PostId,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?)
`
	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return postUser, err
	}

	res, err := statement.Exec(
		postUser.CreatedAt,
		postUser.PostId,
		postUser.UpdatedAt,
		postUser.UserId)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return postUser, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return postUser, err
	}
	log.Println("[sqlite/CreatePostUser] User", postUser.UserId, "can see post", postUser.PostId)
	postUser.PostUserId = int(id)
	return postUser, nil
}

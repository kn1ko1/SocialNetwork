package posts

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Adds post into the given database
func CreatePost(database *sql.DB, post models.Post) (models.Post, error) {

	query := `
	INSERT INTO POSTS (
		Body,
		CreatedAt,
		GroupId,
		ImageURL,
		Privacy,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?, ?, ?, ?)
`
	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing db query.", err)
		return post, err
	}
	res, err := statement.Exec(
		post.Body,
		post.CreatedAt,
		post.GroupId,
		post.ImageURL,
		post.Privacy,
		post.UpdatedAt,
		post.UserId)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return post, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error getting last insert from table.", err)
		return post, err
	}
	post.PostId = int(id)
	return post, nil
}

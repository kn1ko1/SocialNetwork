package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds post into the given database
func CreatePost(database *sql.DB, post models.Post) (models.Post, error) {

	query := `
	INSERT INTO POSTS (
		Body,
		CreatedAt,
		GroupId,
		ImageURL,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?, ?, ?)
`
	statement, err := database.Prepare(query)
	if err != nil {
		return post, err
	}
	res, err := statement.Exec(
		post.Body,
		post.CreatedAt,
		post.GroupId,
		post.ImageURL,
		post.UpdatedAt,
		post.UserId)
	if err != nil {
		return post, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return post, err
	}
	post.PostId = int(id)
	return post, nil
}

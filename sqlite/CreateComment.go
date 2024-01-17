package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Adds comment into the given database
func CreateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {

	query := "INSERT INTO COMMENTS (Body, CreatedAt, ImageURL, PostId, UpdatedAt, UserId) VALUES (?, ?, ?, ?, ?, ?)"
	statement, err := database.Prepare(query)
	if err != nil {
		return comment, err
	}
	res, err := statement.Exec(query, comment.Body, comment.ImageURL, comment.PostId, comment.UserId)
	if err != nil {
		return comment, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.CommentId = int(id)
	return comment, nil
}

package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Adds comment into the given database
func CreateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {

	query := `
	INSERT INTO COMMENTS (
		Body,
		CreatedAt,
		ImageURL,
		PostId,
		UpdatedAt,
		UserId
	) VALUES (?, ?, ?, ?, ?, ?)
`
	statement, err := database.Prepare(query)
	if err != nil {
		return comment, err
	}
	res, err := statement.Exec(
		comment.Body,
		comment.CreatedAt,
		comment.ImageURL,
		comment.PostId,
		comment.UpdatedAt,
		comment.UserId,
	)
	if err != nil {
		utils.HandleError("Error executing statement.", err)
		return comment, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.HandleError("Error retrieving last insert from table.", err)
		return comment, err
	}
	comment.CommentId = int(id)
	return comment, nil
}

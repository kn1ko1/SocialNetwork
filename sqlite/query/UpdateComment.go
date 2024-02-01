package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Updates comment information in the COMMENTS table
func UpdateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {
	query := `
		UPDATE COMMENTS
		SET
			Body = ?,
			ImageURL = ?,
			UpdatedAt = ?
		WHERE CommentId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing statement in UpdateComment.", err)
		return comment, err
	}

	_, err = statement.Exec(
		comment.Body,
		comment.ImageURL,
		comment.UpdatedAt,
		comment.CommentId,
	)

	if err != nil {
		utils.HandleError("Error executing query in UpdateComment.", err)
		return comment, err
	}

	return comment, nil
}

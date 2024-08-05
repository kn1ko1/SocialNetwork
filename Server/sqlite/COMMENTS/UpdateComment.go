package comments

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Updates comment information in the COMMENTS table
func UpdateComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	query := `
		UPDATE COMMENTS
		SET
			Body = (?),
			ImageURL = (?),
			UpdatedAt = (?)
		WHERE CommentId = (?);
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing statement in UpdateComment.", err)
		return comment, err
	}
	_, err = stmt.Exec(
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

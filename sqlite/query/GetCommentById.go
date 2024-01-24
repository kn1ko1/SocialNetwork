package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves comment with the relevant commentId from the COMMENTS table
func GetCommentById(database *sql.DB, commentId int) (models.Comment, error) {
	var comment models.Comment
	err := database.QueryRow("SELECT * FROM COMMENTS WHERE CommentId = ?", commentId).
		Scan(
			&comment.CommentId,
			&comment.Body,
			&comment.CreatedAt,
			&comment.ImageURL,
			&comment.PostId,
			&comment.UpdatedAt,
			&comment.UserId,
		)

	switch {
	case err == sql.ErrNoRows:
		return comment, errors.New("comment not found")
	case err != nil:
		return comment, err
	}

	return comment, nil
}

package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
	"socialnetwork/utils"
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
		utils.HandleError("Comment not found in GetCommentById.", err)
		return comment, errors.New("comment not found")
	case err != nil:
		utils.HandleError("Error retrieving comment by ID in GetCommentById.", err)
		return comment, err
	}

	return comment, nil
}

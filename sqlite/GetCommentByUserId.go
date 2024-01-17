package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves comment with the relevant userId from the COMMENTS table
func GetCommentByUserId(db *sql.DB, userId int) (*models.Comment, error) {
	var comment models.Comment
	err := db.QueryRow("SELECT * FROM COMMENTS WHERE UserId = ?", userId).
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
		return nil, errors.New("comment not found")
	case err != nil:
		return nil, err
	}

	return &comment, nil
}

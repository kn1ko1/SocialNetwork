package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves comment with the relevant groupId from the COMMENTS table
func GetCommentByGroupId(db *sql.DB, groupId int) (*models.Comment, error) {
	var comment models.Comment
	err := db.QueryRow("SELECT * FROM COMMENTS WHERE GroupId = ?", groupId).
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

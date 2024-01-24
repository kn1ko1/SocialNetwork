package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves comment with the relevant userId from the COMMENTS table
func GetCommentsByUserId(database *sql.DB, userId int) ([]models.Comment, error) {
	rows, err := database.Query("SELECT * FROM COMMENTS WHERE UserId = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.CommentId,
			&comment.Body,
			&comment.CreatedAt,
			&comment.ImageURL,
			&comment.PostId,
			&comment.UpdatedAt,
			&comment.UserId,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

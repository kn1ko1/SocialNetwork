package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves comments with the relevant postId from the COMMENTS table
func GetCommentsByPostId(database *sql.DB, postId int) ([]models.Comment, error) {
	rows, err := database.Query("SELECT * FROM COMMENTS WHERE PostId = ?", postId)
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM COMMENTS WHERE PostId = ? statement.", err)
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
			utils.HandleError("Error scanning rows in GetCommentsByPostId.", err)
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetCommentsByPostId.", err)
		return nil, err
	}

	return comments, nil
}

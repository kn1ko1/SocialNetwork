package sqlite

import (
	"database/sql"
	"errors"
	"socialnetwork/models"
)

// Retrieves post with the relevant postId from the POSTS table
func GetPostById(database *sql.DB, groupId int) (*models.Post, error) {
	var post models.Post
	err := database.QueryRow("SELECT * FROM POSTS WHERE PostId = ?", groupId).
		Scan(
			&post.PostId,
			&post.Body,
			&post.CreatedAt,
			&post.GroupId,
			&post.ImageURL,
			&post.UpdatedAt,
			&post.UserId,
		)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("post not found")
	case err != nil:
		return nil, err
	}

	return &post, nil
}

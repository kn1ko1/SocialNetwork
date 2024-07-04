package posts

import (
	"database/sql"
	"errors"
	"socialnetwork/Server/models"
	"socialnetwork/utils"
)

// Retrieves post with the relevant postId from the POSTS table
func GetPostById(database *sql.DB, postId int) (models.Post, error) {
	var post models.Post
	err := database.QueryRow("SELECT * FROM POSTS WHERE PostId = ?", postId).
		Scan(
			&post.PostId,
			&post.Body,
			&post.CreatedAt,
			&post.GroupId,
			&post.ImageURL,
			&post.Privacy,
			&post.UpdatedAt,
			&post.UserId,
		)

	switch {
	case err == sql.ErrNoRows:
		utils.HandleError("Post not found.", err)
		return post, errors.New("post not found")
	case err != nil:
		utils.HandleError("Error querying post by ID.", err)
		return post, err
	}

	return post, nil
}

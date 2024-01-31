package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Retrieves all posts from the POSTS table
func GetAllPosts(database *sql.DB) ([]models.Post, error) {
	rows, err := database.Query("SELECT * FROM POSTS")
	if err != nil {
		utils.HandleError("Error executing SELECT * FROM POSTS statement.", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.PostId,
			&post.Body,
			&post.CreatedAt,
			&post.GroupId,
			&post.ImageURL,
			&post.Privacy,
			&post.UpdatedAt,
			&post.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning rows in GetAllPosts.", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetAllPosts.", err)
		return nil, err
	}

	return posts, nil
}

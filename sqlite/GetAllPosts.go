package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves all posts from the POSTS table
func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query("SELECT * FROM POSTS")
	if err != nil {
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
			&post.UpdatedAt,
			&post.UserId,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

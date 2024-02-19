package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// GetPostsAlmostPrivate retrieves posts for the provided userId from the POST_USERS table
func GetPostsAlmostPrivate(database *sql.DB, userId int) ([]models.Post, error) {
	var posts []models.Post

	// Query to select almost private posts based on the provided userId
	query := `
        SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
        FROM POST_USERS pu
        JOIN POSTS p ON pu.PostId = p.PostId
        WHERE pu.UserId = ?
    `

	rows, err := database.Query(query, userId)
	if err != nil {
		utils.HandleError("Error querying almost private posts by UserId.", err)
		return nil, err
	}
	defer rows.Close()

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
			utils.HandleError("Error scanning row in GetPostsAlmostPrivate.", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsAlmostPrivate.", err)
		return nil, err
	}

	return posts, nil
}

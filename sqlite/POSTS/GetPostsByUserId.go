package posts

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves posts with the relevant userId from the POSTS table
func GetPostsByUserId(database *sql.DB, userId int) ([]models.Post, error) {
	rows, err := database.Query("SELECT * FROM POSTS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error querying posts by UserId.", err)
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
			utils.HandleError("Error scanning row in GetPostsByUserId.", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsByUserId.", err)
		return nil, err
	}

	return posts, nil
}

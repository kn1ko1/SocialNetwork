package posts

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// Retrieves posts with the relevant Privacy from the POSTS table
// Should really only be used to retrieve Public
func GetPostsByPrivacy(database *sql.DB, privacy string) ([]models.Post, error) {
	rows, err := database.Query(
		"SELECT * FROM POSTS WHERE Privacy = ? ORDER BY CreatedAt DESC;",
		privacy)

	if err != nil {
		utils.HandleError("Error querying posts by Privacy.", err)
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
			utils.HandleError("Error scanning row in GetPostsByPrivacy.", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsByPrivacy.", err)
		return nil, err
	}

	return posts, nil
}

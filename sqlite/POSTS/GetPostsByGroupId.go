package posts

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Retrieves posts with the relevant groupId from the POSTS table
func GetPostsByGroupId(database *sql.DB, groupId int) ([]models.Post, error) {
	rows, err := database.Query("SELECT * FROM POSTS WHERE GroupId = ?", groupId)
	if err != nil {
		utils.HandleError("Error querying posts by GroupId.", err)
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
			utils.HandleError("Error scanning row in GetPostsByGroupId.", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsByGroupId.", err)
		return nil, err
	}

	return posts, nil
}

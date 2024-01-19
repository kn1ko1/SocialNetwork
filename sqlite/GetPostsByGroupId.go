package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Retrieves posts with the relevant groupId from the POSTS table
func GetPostsByGroupId(database *sql.DB, groupId int) ([]*models.Post, error) {
	rows, err := database.Query("SELECT * FROM POSTS WHERE GroupId = ?", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post

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

		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

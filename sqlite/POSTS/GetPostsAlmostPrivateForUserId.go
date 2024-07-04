package posts

import (
	"database/sql"
	"socialnetwork/Server/models"
	"socialnetwork/Server/utils"
)

// GetPostsAlmostPrivateWithComments retrieves almost private posts for the provided userId along with associated comments
func GetPostsAlmostPrivateForUserId(database *sql.DB, userId int) ([]models.Post, error) {
	var result []models.Post

	// Query to select almost private posts based on the provided userId
	query := `
        SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
        FROM POST_USERS pu
        JOIN POSTS p ON pu.PostId = p.PostId
        WHERE pu.UserId = ?
		ORDER BY 
    	p.CreatedAt DESC;
    `

	rows, err := database.Query(query, userId)
	if err != nil {
		// no results found in DB
		return result, nil
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
	}

	return result, nil
}

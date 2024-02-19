package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// GetPrivatePostsForFollower retrieves private posts for the given followerId
func GetPostsPrivate(database *sql.DB, userId int) ([]models.Post, error) {
	var posts []models.Post

	query := `
        SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
        FROM POSTS p
        JOIN USER_USERS uu ON uu.SubjectId = p.UserId
        WHERE uu.FollowerId = ? AND p.Privacy = 'private'
    `

	rows, err := database.Query(query, userId)
	if err != nil {
		utils.HandleError("Error querying private posts for follower.", err)
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
			utils.HandleError("Error scanning row in GetPostsPrivate.", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsPrivate.", err)
		return nil, err
	}

	return posts, nil
}

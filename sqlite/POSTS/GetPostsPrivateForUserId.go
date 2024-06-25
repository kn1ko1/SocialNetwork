package posts

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// GetPostsPrivateForUserId retrieves private posts for the given followerId along with associated comments
func GetPostsPrivateForUserId(database *sql.DB, userId int) ([]models.Post, error) {
	var result []models.Post

	query := `
	SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
	FROM POSTS p
	WHERE p.Privacy = 'private' AND (
		p.UserId = ? OR p.UserId IN (
			SELECT uu.SubjectId
			FROM USER_USERS uu
			WHERE uu.FollowerId = ?
		)
	)
	ORDER BY p.CreatedAt DESC;
    `

	rows, err := database.Query(query, userId, userId)
	if err != nil {
		// no entries found in DB
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
			utils.HandleError("Error scanning row in GetPostsPrivate.", err)
			return nil, err
		}
	}
	return result, nil
}

package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// GetPostsAlmostPrivateWithComments retrieves almost private posts for the provided userId along with associated comments
func GetPostsAlmostPrivateWithComments(database *sql.DB, userId int) ([]transport.PostWithComments, error) {
	var result []transport.PostWithComments

	// Query to select almost private posts based on the provided userId
	query := `
        SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
        FROM POST_USERS pu
        JOIN POSTS p ON pu.PostId = p.PostId
        WHERE pu.UserId = ?
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

		// Get comments associated with the current post
		comments, err := GetCommentsByPostId(database, post.PostId)
		if err != nil {
			utils.HandleError("Error getting comments for post.", err)
			return nil, err
		}

		// Append the post along with its comments to the result
		postWithComments := transport.PostWithComments{
			Post:     post,
			Comments: comments,
		}
		result = append(result, postWithComments)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsAlmostPrivate.", err)
		return nil, err
	}

	return result, nil
}

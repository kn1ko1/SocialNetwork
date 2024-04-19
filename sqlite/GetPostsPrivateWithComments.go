package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	comments "socialnetwork/sqlite/COMMENTS"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// GetPostsPrivateWithComments retrieves private posts for the given followerId along with associated comments
func GetPostsPrivateWithComments(database *sql.DB, userId int) ([]transport.PostWithComments, error) {
	var result []transport.PostWithComments

	query := `
        SELECT p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt, p.UserId
        FROM POSTS p
        JOIN USER_USERS uu ON uu.SubjectId = p.UserId
        WHERE uu.FollowerId = ? AND p.Privacy = 'private'
		ORDER BY 
    	p.CreatedAt DESC;
    `

	rows, err := database.Query(query, userId)
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

		// Get comments associated with the current post
		comments, err := comments.GetCommentsByPostId(database, post.PostId)
		if err != nil {
			utils.HandleError("Error getting comments for post.", err)
			return nil, err
		}

		// Append the post along with its comments to the result
		result = append(result, transport.PostWithComments{
			Post:     post,
			Comments: comments,
		})
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetPostsPrivate.", err)
		return nil, err
	}

	return result, nil
}

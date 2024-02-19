package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves public (group 0) posts and relevant comments
func GetPublicPostsWithComments(database *sql.DB) ([]transport.PostWithComments, error) {
	// Query to fetch posts and their corresponding comments using foreign keys
	query := `
SELECT 
	   p.PostId, p.Body, p.CreatedAt, p.GroupId, p.ImageURL, p.Privacy, p.UpdatedAt,
	c.CommentId, c.Body, c.CreatedAt, c.ImageURL, c.UpdatedAt, c.UserId
FROM 
	   POSTS p
LEFT JOIN 
	COMMENTS c ON p.PostId = c.PostId
WHERE 
	p.GroupId = 0

`

	rows, err := database.Query(query)
	if err != nil {
		utils.HandleError("Error querying posts by UserId.", err)
		return nil, err
	}
	defer rows.Close()

	// Map to store posts with their comments
	postMap := make(map[int]*transport.PostWithComments)

	for rows.Next() {
		var postWithComment transport.PostWithComments
		var post models.Post
		var comment models.Comment

		err := rows.Scan(
			&post.PostId, &post.Body, &post.CreatedAt, &post.GroupId, &post.ImageURL, &post.Privacy, &post.UpdatedAt,
			&comment.CommentId, &comment.Body, &comment.CreatedAt, &comment.ImageURL, &comment.UpdatedAt, &comment.UserId,
		)
		if err != nil {
			utils.HandleError("Error scanning row in GetHomeDataForUser.", err)
			continue
		}

		if _, ok := postMap[post.PostId]; !ok {
			postWithComment.Post = post
			postWithComment.Comments = []models.Comment{}
			postMap[post.PostId] = &postWithComment
		}

		postMap[post.PostId].Comments = append(postMap[post.PostId].Comments, comment)
	}

	if err := rows.Err(); err != nil {
		utils.HandleError("Error iterating over rows in GetHomeDataForUser.", err)
		return nil, err
	}

	var result []transport.PostWithComments
	for _, postWithComment := range postMap {
		result = append(result, *postWithComment)
	}

	return result, nil
}

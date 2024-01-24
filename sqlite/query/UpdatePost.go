package sqlite

import (
	"database/sql"
	"socialnetwork/models"
)

// Updates post information in the POSTS table
func UpdatePost(database *sql.DB, post models.Post) (models.Post, error) {
	query := `
		UPDATE POSTS
		SET
			Body = ?,
			ImageUrl = ?,
			UpdatedAt = ?
		WHERE PostId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		return post, err
	}

	_, err = statement.Exec(
		post.Body,
		post.ImageURL,
		post.UpdatedAt,
		post.PostId,
	)

	if err != nil {
		return post, err
	}

	return post, nil
}

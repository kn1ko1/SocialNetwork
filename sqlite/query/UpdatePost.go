package sqlite

import (
	"database/sql"
	utils "socialnetwork/helper"
	"socialnetwork/models"
)

// Updates post information in the POSTS table
func UpdatePost(database *sql.DB, post models.Post) (models.Post, error) {
	query := `
		UPDATE POSTS
		SET
			Body = ?,
			ImageUrl = ?,
			UpdatedAt = ?,
			Privacy = ?
		WHERE PostId = ?
	`

	statement, err := database.Prepare(query)
	if err != nil {
		utils.HandleError("Error preparing database in UpdatePost.", err)

		return post, err
	}

	_, err = statement.Exec(
		post.Body,
		post.ImageURL,
		post.UpdatedAt,
		post.Privacy,
		post.PostId,
	)

	if err != nil {
		utils.HandleError("Error executing statement in UpdatePost.", err)

		return post, err
	}

	return post, nil
}

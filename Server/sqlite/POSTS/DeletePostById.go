package posts

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// deletes a specific post from the POSTS table
func DeletePostById(database *sql.DB, postId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE PostId = ?", postId)
	if err != nil {
		utils.HandleError("Error executing delete post by ID statement.", err)
		return err
	}
	return nil
}

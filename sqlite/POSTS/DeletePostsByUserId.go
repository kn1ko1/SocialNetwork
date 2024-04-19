package posts

import (
	"database/sql"
	"socialnetwork/utils"
)

// deletes posts related to userId from the POSTS table
func DeletePostsByUserId(database *sql.DB, userId int) error {
	_, err := database.Exec("DELETE FROM POSTS WHERE UserId = ?", userId)
	if err != nil {
		utils.HandleError("Error executing delete posts by userId statement.", err)
		return err
	}
	return nil
}

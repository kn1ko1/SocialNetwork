package user_users

import (
	"database/sql"
	"socialnetwork/Server/utils"
)

// deletes user users related to followerId from the USER_USERS table
func DeleteUserUsersByFollowerId(database *sql.DB, followerId int) error {
	_, err := database.Exec("DELETE FROM USER_USERS WHERE FollowerId = ?", followerId)
	if err != nil {
		utils.HandleError("Error executing delete user users by follower Id.", err)
		return err
	}
	return nil
}

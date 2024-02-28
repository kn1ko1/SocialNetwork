package sqlite

import (
	"database/sql"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves data for the user's homepage including posts and comments
func GetGroupDataForUser(database *sql.DB, userId int) (transport.HomeModel, error) {

	var userHomeData transport.HomeModel
	var err error

	// Get all users
	userHomeData.AllUsers, err = GetAllUsers(database)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// Get public posts with comments
	userHomeData.PublicPostsWithComments, err = GetPublicPostsWithComments(database)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// GetPostsPrivate retrieves private posts for the given followerId
	userHomeData.PrivatePosts, err = GetPostsPrivateWithComments(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// GetPostsAlmostPrivate retrieves posts for the provided userId from the POST_USERS table

	userHomeData.AlmostPrivatePosts, err = GetPostsAlmostPrivateWithComments(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// Get groups that the user is a member of
	userHomeData.UserGroups, err = GetGroupsByUserId(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// Get events associated with the user
	userHomeData.UserEvents, err = GetEventsByUserId(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	// Get notifications for the user
	userHomeData.UserNotifications, err = GetNotificationsByUserId(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	return userHomeData, nil
}

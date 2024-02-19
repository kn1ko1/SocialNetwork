package sqlite

import (
	"database/sql"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves data for the user's homepage including posts and comments
func GetHomeDataForUser(database *sql.DB, userId int) (transport.HomeModel, error) {

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

	userHomeData.PrivatePosts, err = GetPostsPrivate(database, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		return userHomeData, err
	}

	userHomeData.AlmostPrivatePosts, err = GetPostsAlmostPrivate(database, userId)
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

// Retrieves data for the users homepage including posts and comments
// func GetHomeDataForUser(database *sql.DB, userId int) ([]models.Post, []models.Comment, error) {
// 	var PostsWithComments []PostWithComments
// 	publicPosts, _ := GetPostsByGroupId(database, 0)
// 	// publicPostsCommentsw := ... err
// 	// postsByFollowing := ... err
// 	// postsByFollowerByChosen := ... err

// 	allUsers, _ := GetAllUsers(database)

// 	userGroups, _ := GetGroupUsersByUserId(database, userId)
// 	userEvents, _ := GetEventsByUserId(database, userId)
// 	userNotifications, _ := GetNotificationsByUserId(database, userId)
// }

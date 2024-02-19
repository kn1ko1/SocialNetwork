package sqlite

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/transport"
)

// Retrieves data for the user's homepage including posts and comments
func GetHomeDataForUser(database *sql.DB, userId int) ([]models.User, []transport.PostWithComments, []models.Group, []models.Event, []models.Notification, error) {
	var (
		allUsers                []models.User
		publicPostsWithComments []transport.PostWithComments
		userGroups              []models.Group
		userEvents              []models.Event
		userNotifications       []models.Notification
		err                     error
	)

	// Get all users
	allUsers, err = GetAllUsers(database)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Get public posts with comments
	publicPostsWithComments, err = GetPublicPostsWithComments(database)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Get groups that the user is a member of
	userGroups, err = GetGroupsByUserId(database, userId)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Get events associated with the user
	userEvents, err = GetEventsByUserId(database, userId)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Get notifications for the user
	userNotifications, err = GetNotificationsByUserId(database, userId)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return allUsers, publicPostsWithComments, userGroups, userEvents, userNotifications, nil
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

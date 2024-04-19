package sqlite

import (
	"database/sql"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Needs to be split - not a SQLite function
// Retrieves data for the user's homepage including posts and comments
func GetHomeDataForUser(identityDB, businessDb *sql.DB, userId int) (transport.HomeModel, error) {

	var userHomeData transport.HomeModel
	var err error

	// Get public posts with comments
	userHomeData.AllUsers, err = GetAllUsersTransport(identityDB)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		// return userHomeData, err
	}

	followedUsers, err := GetUserUsersByFollowerId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error getting followedUsers in GetHomeDataForUser", err)
		// return userHomeData, err
	}
	for i := range userHomeData.AllUsers {
		for _, followedUser := range followedUsers {
			if userHomeData.AllUsers[i].UserId == followedUser.SubjectId {
				userHomeData.AllUsers[i].IsFollowed = true
				break
			}
		}
	}
	// Get public posts with comments
	userHomeData.PublicPostsWithComments, err = GetPublicPostsWithComments(businessDb)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		// return userHomeData, err
	}

	// GetPostsPrivate retrieves private posts for the given followerId
	userHomeData.PrivatePosts, err = GetPostsPrivateWithComments(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		// return userHomeData, err
	}

	// GetPostsAlmostPrivate retrieves posts for the provided userId from the POST_USERS table

	userHomeData.AlmostPrivatePosts, err = GetPostsAlmostPrivateWithComments(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		// return userHomeData, err
	}

	// Get groups that the user is a member of
	userHomeData.UserGroups, err = GetGroupsByUserId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetHomeDataForUser", err)
		// return userHomeData, err
	}

	return userHomeData, nil
}

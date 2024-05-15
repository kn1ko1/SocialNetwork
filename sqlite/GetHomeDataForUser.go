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

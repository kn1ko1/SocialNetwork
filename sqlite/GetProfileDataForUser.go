package sqlite

import (
	"database/sql"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves data for the user's homepage including posts and comments
func GetProfileDataForUser(businessDb *sql.DB, userId int) (transport.ProfileModel, error) {

	var userProfileData transport.ProfileModel
	var err error

	userProfileData.UserPostData, err = GetPostsByUserId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}

	userProfileData.UserFollowerData, err = GetUserUsersBySubjectId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}

	userProfileData.UserFollowsData, err = GetUserUsersByFollowerId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}

	return userProfileData, nil
}

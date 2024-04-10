package sqlite

import (
	"database/sql"
	"log"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Retrieves data for the user's homepage including posts and comments
func GetProfileDataForUser(identityDb *sql.DB, businessDb *sql.DB, userId int) (transport.ProfileModel, error) {

	var userProfileData transport.ProfileModel
	var err error
	userData, err := GetUserById(identityDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}
	userProfileData.ProfileUserData = transport.ProfileRegistrationInfo{
		UserId:    userData.UserId,
		Bio:       userData.Bio,
		DOB:       userData.DOB,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		ImageURL:  userData.ImageURL,
		IsPublic:  userData.IsPublic,
		LastName:  userData.LastName,
		Username:  userData.Username,
	}
	userProfileData.UserPostData, err = GetPostsByUserId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}

	followerUserUsers, err := GetUserUsersBySubjectId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}
	log.Println("len(followerUserUsers)", len(followerUserUsers))

	// var userFollowersData []transport.UserTransport
	for i := 0; i < len(followerUserUsers); i++ {
		userFollowerData, err := GetUsernameByUserId(identityDb, followerUserUsers[i].FollowerId)
		if err != nil {
			utils.HandleError("Error in GetProfileDataForUser", err)
		}
		userProfileData.UserFollowerData = append(userProfileData.UserFollowerData, userFollowerData)
	}

	followsUsersUsers, err := GetUserUsersByFollowerId(businessDb, userId)
	if err != nil {
		utils.HandleError("Error in GetProfileDataForUser", err)
		// return userProfileData, err
	}
	log.Println("len(followsUsersUsers)", len(followsUsersUsers))
	//var userFollowsData []transport.UserTransport
	for i := 0; i < len(followsUsersUsers); i++ {
		userFollowData, err := GetUsernameByUserId(identityDb, followsUsersUsers[i].SubjectId)
		if err != nil {
			utils.HandleError("Error in GetProfileDataForUser", err)
		}
		userProfileData.UserFollowsData = append(userProfileData.UserFollowsData, userFollowData)
	}

	log.Println("this is userProfileData.UserFollowerData", userProfileData.UserFollowerData)
	log.Println("this is userProfileData.UserFollowsData", userProfileData.UserFollowsData)
	return userProfileData, nil
}

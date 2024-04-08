package transport

import "socialnetwork/models"

type ProfileModel struct {
	ProfileUserData  ProfileRegistrationInfo `json:"profileUserData"`
	UserPostData     []models.Post           `json:"userPostData"`
	UserFollowerData []models.UserUser       `json:"userFollowerData"`
	UserFollowsData  []models.UserUser       `json:"userFollowsData"`
}

package transport

import "socialnetwork/Server/models"

type ProfileModel struct {
	ProfileUserData  ProfileRegistrationInfo `json:"profileUserData"`
	UserPostData     []models.Post           `json:"userPostData"`
	UserFollowerData []models.User           `json:"userFollowerData"`
	UserFollowsData  []models.User           `json:"userFollowsData"`
}

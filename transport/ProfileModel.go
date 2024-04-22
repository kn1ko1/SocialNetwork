package transport

import "socialnetwork/models"

type ProfileModel struct {
	ProfileUserData  ProfileRegistrationInfo `json:"profileUserData"`
	UserPostData     []models.Post           `json:"userPostData"`
	UserFollowerData []UserTransport         `json:"userFollowerData"`
	UserFollowsData  []UserTransport         `json:"userFollowsData"`
}

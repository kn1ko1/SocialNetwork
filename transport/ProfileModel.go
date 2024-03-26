package transport

import "socialnetwork/models"

type ProfileModel struct {
	ProfileUserData  ProfileRegistrationInfo `json:"profileUserData"`
	UserPostData     []models.Post           `json:"userPostData"`
	UserFollowerData []models.UserUser       `json:"userFollowerData"`
	UserFollowsData  []models.UserUser       `json:"userFollowsData"`
}

type ProfileRegistrationInfo struct {
	UserId    int    `json:"userId"`
	Bio       string `json:"bio"`
	DOB       int64  `json:"dob"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	ImageURL  string `json:"imageURL"`
	IsPublic  bool   `json:"isPublic"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

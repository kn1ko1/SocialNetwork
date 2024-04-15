package transport

import "socialnetwork/models"

type HomeModel struct {
	AllUsers                []UserTransport    `json:"userList"`
	AlmostPrivatePosts      []PostWithComments `json:"almostPrivatePosts"`
	PrivatePosts            []PostWithComments `json:"privatePosts"`
	PublicPostsWithComments []PostWithComments `json:"publicPostsWithComments"`
	UserGroups              []models.Group     `json:"userGroups"`
}

package transport

import "socialnetwork/models"

type HomeModel struct {
	AlmostPrivatePosts      []PostWithComments `json:"almostPrivatePosts"`
	PrivatePosts            []PostWithComments `json:"privatePosts"`
	PublicPostsWithComments []PostWithComments `json:"publicPostsWithComments"`
	UserGroups              []models.Group     `json:"userGroups"`
}

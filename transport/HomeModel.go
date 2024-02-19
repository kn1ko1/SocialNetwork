package transport

import "socialnetwork/models"

type HomeModel struct {
	PostsWithComments       []PostWithComments    `json:"postsWithComments"`
	AllUsers                []models.User         `json:"users"`
	PublicPostsWithComments []PostWithComments    `json:"publicPostsWithComments"`
	PrivatePosts            []models.Post         `json:"privatePosts"`
	AlmostPrivatePosts      []models.Post         `json:"almostPrivatePosts"`
	UserGroups              []models.Group        `json:"userGroups"`
	UserEvents              []models.Event        `json:"userEvents"`
	UserNotifications       []models.Notification `json:"userNotifications"`
}

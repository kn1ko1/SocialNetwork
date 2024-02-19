package transport

import "socialnetwork/models"

type HomeModel struct {
	AllUsers                []models.User         `json:"users"`
	AlmostPrivatePosts      []models.Post         `json:"almostPrivatePosts"`
	PrivatePosts            []models.Post         `json:"privatePosts"`
	PublicPostsWithComments []PostWithComments    `json:"publicPostsWithComments"`
	UserEvents              []models.Event        `json:"userEvents"`
	UserGroups              []models.Group        `json:"userGroups"`
	UserNotifications       []models.Notification `json:"userNotifications"`
}

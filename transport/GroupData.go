package transport

import "socialnetwork/models"

type GroupData struct {
	AllUsers                []UserTransport
	GroupUsersWithUsernames []UserTransport
	GroupPosts              []models.Post
	GroupMessages           []models.Message
	GroupEvents             []models.Event
}

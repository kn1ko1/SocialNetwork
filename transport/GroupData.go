package transport

import "socialnetwork/models"

type GroupData struct {
	AllUsers      []UserTransport    `json:"allUsers"`
	GroupUsers    []models.GroupUser `json:"groupUsers"`
	GroupPosts    []models.Post      `json:"groupPosts"`
	GroupMessages []models.Message   `json:"groupMessages"`
	GroupEvents   []models.Event     `json:"groupEvents"`
}

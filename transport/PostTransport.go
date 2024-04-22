package transport

import "socialnetwork/models"

type PostTransport struct {
	models.Post
	Username string `json:"username"`
}

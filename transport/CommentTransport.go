package transport

import "socialnetwork/models"

type CommentTransport struct {
	models.Comment
	Username string `json:"username"`
}

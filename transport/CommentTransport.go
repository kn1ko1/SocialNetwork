package transport

import "socialnetwork/models"

type CommentTransport struct {
	Comment models.Comment `json:"comment"`
	User    models.User    `json:"user"`
}

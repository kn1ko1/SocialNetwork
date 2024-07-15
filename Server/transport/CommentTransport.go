package transport

import "socialnetwork/Server/models"

type CommentTransport struct {
	Comment models.Comment `json:"comment"`
	User    models.User    `json:"user"`
}

package transport

import "socialnetwork/models"

type PostTransport struct {
	Post models.Post `json:"post"`
	User models.User `json:"user"`
}

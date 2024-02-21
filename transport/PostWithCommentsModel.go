package transport

import "socialnetwork/models"

type PostWithComments struct {
	Comments []models.Comment `json:"comments"`
	Post     models.Post      `json:"post"`
}

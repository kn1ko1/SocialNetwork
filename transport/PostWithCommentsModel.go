package transport

import "socialnetwork/models"

type PostWithComments struct {
	Post     models.Post      `json:"post"`
	Comments []models.Comment `json:"comments"`
}

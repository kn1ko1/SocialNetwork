package repo

import "socialnetwork/models"

type IRepository interface {
	CreatePost(p models.Post) (models.Post, error)
}

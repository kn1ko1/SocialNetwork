package repo

import "matthewhope/example-architecture/models"

type IRepository interface {
	CreatePost(p models.Post) (models.Post, error)
}

package repo

import "socialnetwork/models"

type IPostUsers interface {
	// PostUser
	CreatePostUser(postUser models.PostUser) (models.PostUser, error)
	GetPostUsersByUserId(userId int) ([]models.PostUser, error)
	GetPostUsersByPostId(postId int) ([]models.PostUser, error)
	DeletePostUsersByUserId(userId int) error
	DeletePostUsersByPostId(postId int) error
	DeletePostUserByPostIdAndUserId(postId, userId int) error
	DeleteAllPostUsers() error
}

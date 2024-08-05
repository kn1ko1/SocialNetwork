package repo

import "socialnetwork/Server/models"

type IPosts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostById(postId int) (models.Post, error)
	GetPostsByGroupId(groupId int) ([]models.Post, error)
	GetPostsByUserId(userId int) ([]models.Post, error)
	GetPostsAlmostPrivateForUserId(userId int) ([]models.Post, error)
	GetPostsPrivateForUserId(userId int) ([]models.Post, error)
	GetPostsByPrivacy(privacy string) ([]models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeletePostById(postId int) error
	DeletePostByGroupId(groupId int) error
	DeletePostsByUserId(userId int) error
	DeleteAllPosts() error
}

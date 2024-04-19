package repo

import "socialnetwork/models"

type IComments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
	GetAllComments() ([]models.Comment, error)
	GetCommentById(commentId int) (models.Comment, error)
	// GetCommentsByGroupId(groupId int) ([]models.Comment, error)
	GetCommentsByUserId(userId int) ([]models.Comment, error)
	GetCommentsByPostId(postId int) ([]models.Comment, error)
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteCommentById(commentId int) error
	DeleteCommentsByGroupId(groupId int) error
	DeleteCommentsByUserId(userId int) error
	DeleteCommentsByPostId(postId int) error
	DeleteAllComments() error
}

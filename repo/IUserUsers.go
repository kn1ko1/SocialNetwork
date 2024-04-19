package repo

import "socialnetwork/models"

type IUserUsers interface {
	//UserUsers ... Yes, I know
	CreateUserUser(userUser models.UserUser) (models.UserUser, error)
	GetUserUsersBySubjectId(subjectId int) ([]models.UserUser, error)
	GetUserUsersByFollowerId(followerId int) ([]models.UserUser, error)
	GetUserUserByFollowerIdAndSubjectId(followerId, subjectId int) (models.UserUser, error)
	DeleteUserUsersBySubjectId(subjectId int) error
	DeleteUserUsersByFollowerId(followerId int) error
	DeleteUserUserBySubjectIdAndFollowerId(subjectId, followerId int) error
}

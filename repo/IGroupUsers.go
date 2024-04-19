package repo

import "socialnetwork/models"

type IGroupUsers interface {
	//GroupUser
	CreateGroupUser(groupUser models.GroupUser) (models.GroupUser, error)
	GetGroupUser(GroupUserId int) (models.GroupUser, error)
	GetGroupUsersByUserId(userId int) ([]models.GroupUser, error)
	GetGroupUsersByGroupId(groupId int) ([]models.GroupUser, error)
	DeleteGroupUsersByUserId(UserId int) error
	DeleteGroupUserByGroupId(groupId int) error
	DeleteGroupUserByGroupIdAndUserId(groupId, userId int) error
	DeleteGroupUser(groupUser int) error
	DeleteAllGroupUsers() error
}

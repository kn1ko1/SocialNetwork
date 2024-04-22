package repo

import "socialnetwork/models"

type IGroups interface {
	//Group
	CreateGroup(group models.Group) (models.Group, error)
	GetGroupById(groupId int) (models.Group, error)
	GetAllGroups() ([]models.Group, error)
	UpdateGroup(group models.Group) (models.Group, error)
	DeleteGroup(groupId int) error
	DeleteAllGroups() error
}

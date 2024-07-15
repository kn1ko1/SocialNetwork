package repo

import "socialnetwork/Server/models"

type IEvents interface {
	CreateEvent(event models.Event) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
	GetEventById(eventId int) (models.Event, error)
	GetEventsByGroupId(groupId int) ([]models.Event, error)
	GetEventsByUserId(userId int) ([]models.Event, error)
	UpdateEvent(event models.Event) (models.Event, error)
	DeleteEventById(eventId int) error
	DeleteEventsByGroupId(groupId int) error
	DeleteEventsByUserId(userId int) error
	DeleteAllEvents() error
}

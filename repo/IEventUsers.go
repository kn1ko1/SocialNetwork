package repo

import "socialnetwork/models"

type IEventUsers interface {
	//EventUser
	CreateEventUser(event models.EventUser) (models.EventUser, error)
	GetEventUsersByUserId(userId int) ([]models.EventUser, error)
	GetEventUsersByEventId(eventId int) ([]models.EventUser, error)
	DeleteEventUsersByUserId(userId int) error
	DeleteEventUsersByEventId(eventId int) error
	DeleteEventUserByEventIdAndUserId(eventId, userId int) error
	DeleteAllEventUsers() error
}

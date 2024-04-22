package repo

import "socialnetwork/models"

type IMessages interface {
	// Message
	CreateMessage(message models.Message) (models.Message, error)
	// GetAllMessages() ([]models.Message, error)
	GetMessagesByMessageTypeandTargetId(messageType string, targetId int) ([]models.Message, error)
	GetMessageById(messageId int) (models.Message, error)
	GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error)
	UpdateMessage(message models.Message) (models.Message, error)
	// DeleteMessagesByType(messageType string) error
	DeleteMessageById(messageId int) error
	DeleteMessagesBySenderId(senderId int) error
	// DeleteMessagesByTargetId(targetId int) error
	DeleteAllMessages() error
}

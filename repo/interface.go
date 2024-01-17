package repo

import "socialnetwork/models"

type IRepository interface {
	// User
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	DeleteAllUsers() error
	DeleteUserById(userId int) error

	// Post
	CreatePost(post models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostById(postId int) ([]models.Post, error)
	GetPostsByGroupId(groupId int) (models.Post, error)
	GetPostsByUserId(userId int) (models.Post, error)
	DeleteAllPosts() error
	DeletePostById(postId int) error
	DeletePostByGroupId(groupId int) error
	DeletePostByUserId(userId int) error

	// Comments
	CreateComment(comment models.Comment) error
	GetAllComments() ([]models.Comment, error)
	GetCommentById(commentId int) (models.Comment, error)
	GetCommentsByGroupId(groupId int) (models.Comment, error)
	GetCommentsByUserId(userId int) (models.Comment, error)
	GetCommentsByPostId(postId int) (models.Comment, error)
	DeleteAllComments() error
	DeleteCommentById(commentId int) error
	DeleteCommentsByGroupId(groupId int) error
	DeleteCommentsByUserId(userId int) error
	DeleteCommentsByPostId(postId int) error

	// Event
	CreateEvent(event models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetEventById(eventId int) (models.Event, error)
	GetEventsByGroupId(groupId int) (models.Event, error)
	GetEventsByUserId(userId int) (models.Event, error)
	DeleteAllEvents() error
	DeleteEventById(eventId int) error
	DeleteEventsByGroupId(groupId int) error
	DeleteEventsByUserId(userId int) error

	// Message
	CreateMessage(message models.Message) error
	GetAllMessages() ([]models.Message, error)
	GetMessagesByType(messageType string) ([]models.Message, error)
	GetMessageById(messageId int) (models.Message, error)
	GetMessagesBySenderId(senderId int) (models.Message, error)
	GetMessagesByTargetId(targetId int) (models.Message, error)
	DeleteAllMessages() error
	DeleteMessagesByType(messageType string) error
	DeleteMessageById(messageId int) error
	DeleteMessagesBySenderId(senderId int) error
	DeleteMessagesByTargetId(targetId int) error
}

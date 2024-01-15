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
	DeleteUserByEmail(email string) error
	DeleteUserByUsername(username string) error

	// Post
	CreatePost(post models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostById(postId int) ([]models.Post, error)
	GetPostByGroupId(groupId int) (models.Post, error)
	GetPostByUserId(userId int) (models.Post, error)
	DeleteAllPosts() error
	DeletePostById(postId int) error
	DeletePostByGroupId(groupId int) error
	DeletePostByUserId(userId int) error

	// Comments
	CreateComment(comment models.Comment) error
	GetAllComments() ([]models.Comment, error)
	GetCommentById(commentId int) (models.Comment, error)
	GetCommentByGroupId(groupId int) (models.Comment, error)
	GetCommentByUserId(userId int) (models.Comment, error)
	GetCommentByPostId(postId int) (models.Comment, error)
	DeleteAllComment() error
	DeleteCommentById(commentId int) error
	DeleteCommentByGroupId(groupId int) error
	DeleteCommentByUserId(userId int) error
	DeleteCommentByPostId(postId int) error

	// Event
	CreateEvent(event models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetEventById(eventId int) (models.Event, error)
	GetEventByGroupId(groupId int) (models.Event, error)
	GetEventByUserId(userId int) (models.Event, error)
	DeleteAllEvent() error
	DeleteEventById(eventId int) error
	DeleteEventByGroupId(groupId int) error
	DeleteEventByUserId(userId int) error

	// Message
	CreateMessage(message models.Message) error
	GetAllMessages() ([]models.Message, error)
	GetAllMessagesByType(messageType string) ([]models.Message, error)
	GetMessageById(messageId int) (models.Message, error)
	GetMessageBySenderId(senderId int) (models.Message, error)
	GetMessageByTargetId(targetId int) (models.Message, error)
	DeleteAllMessage() error
	DeleteAllMessagesByType(messageType string) error
	DeleteMessageById(messageId int) error
	DeleteMessageBySenderId(senderId int) error
	DeleteMessageByTargetId(targetId int) error
}

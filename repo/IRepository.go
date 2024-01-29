// Message from Matt
// Please do not touch this file

package repo

import "socialnetwork/models"

type IRepository interface {
	// User
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUserById(userId int) error
	DeleteAllUsers() error

	// Post
	CreatePost(post models.Post) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostById(postId int) (models.Post, error)
	GetPostsByGroupId(groupId int) ([]models.Post, error)
	GetPostsByUserId(userId int) ([]models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeletePostById(postId int) error
	DeletePostByGroupId(groupId int) error
	DeletePostsByUserId(userId int) error
	DeleteAllPosts() error

	// Comments
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

	// Event
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

	// Message
	CreateMessage(message models.Message) (models.Message, error)
	// GetAllMessages() ([]models.Message, error)
	// GetMessagesByType(messageType string) ([]models.Message, error)
	GetMessageById(messageId int) (models.Message, error)
	GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error)
	UpdateMessage(message models.Message) (models.Message, error)
	// DeleteMessagesByType(messageType string) error
	DeleteMessageById(messageId int) error
	DeleteMessagesBySenderId(senderId int) error
	// DeleteMessagesByTargetId(targetId int) error
	DeleteAllMessages() error

	//Group
	CreateGroup(group models.Group) (models.Group, error)
	GetAllGroups() ([]models.Group, error)
	UpdateGroup(group models.Group) (models.Group, error)
	DeleteAllGroups() error

	//Group_User
	CreateGroup_User(groupUser models.GroupUser) (models.GroupUser, error)
	GetGroup_UserByUserId(userId int) (models.GroupUser, error)

	//Notification
	CreateNotification(notification models.Notification) (models.Notification, error)
	GetNotificationById(notificationId int) (models.Notification, error)
	UpdateNotification(notification models.Notification) (models.Notification, error)
	DeleteNotificationById(notificationId int) error
}

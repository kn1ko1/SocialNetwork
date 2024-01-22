package repo

import (
	"database/sql"
	"socialnetwork/models"
)

type SQLiteRepository struct {
}

func NewSQLiteRepository() *SQLiteRepository {
	return &SQLiteRepository{}
}

func (r *SQLiteRepository) CreateUser(CreatedAt, DOB date, Email string, EncryptedPassword string, FirstName string, LastName string, Username string) (int, error) {
	return sql.CreateUser(postID, authorID, author, content, timestamp)
}

func (r *SQLiteRepository) CreateMessage(senderID, targetID int, author, content string, timestamp int64) (int, error) {
	return sql.CreateMessage(senderID, targetID, author, content, timestamp)
}

func (r *SQLiteRepository) CreatePost(authorID int, author, content, categories string, timestamp int64) (int, error) {
	return sql.CreatePost(authorID, author, content, categories, timestamp)
}

func (r *SQLiteRepository) CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password string) (int, error) {
	return sql.CreateUser(nickname, age, gender, firstName, lastName, emailAddress, password)
}

func (r *SQLiteRepository) GetUserByID(id int) (models.User, error) {
	return sql.GetUserByID(id)
}

func (r *SQLiteRepository) GetUsers() ([]models.User, error) {
	return sql.GetAllUsers()
}

func (r *SQLiteRepository) GetUserByNickname(nickname string) (models.User, error) {
	return sql.GetUserByNickname(nickname)
}

func (r *SQLiteRepository) GetMessagesBySenderAndTargetIDs(senderID, targetID int) ([]models.Message, error) {
	return sql.GetMessagesBySenderAndTargetIDs(senderID, targetID)
}

func (r *SQLiteRepository) GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset int) ([]models.Message, error) {
	return sql.GetLimitedMessagesBySenderAndTargetIDs(senderID, targetID, limit, offset)
}

func (r *SQLiteRepository) GetPosts() ([]models.Post, error) {
	return sql.GetAllPosts()
}

func (r *SQLiteRepository) GetCommentsByPostID(postID int) ([]models.Comment, error) {
	return sql.GetCommentsByPostID(postID)
}

type IRepy interface {
	// User
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteAllUsers() error
	DeleteUserById(userId int) error

	// Post
	CreatePost(post models.Post) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostById(postId int) ([]models.Post, error)
	GetPostsByGroupId(groupId int) (models.Post, error)
	GetPostsByUserId(userId int) (models.Post, error)
	UpdatePost(post models.Post) (models.Post, error)
	DeleteAllPosts() error
	DeletePostById(postId int) error
	DeletePostByGroupId(groupId int) error
	DeletePostByUserId(userId int) error

	// Comments
	CreateComment(comment models.Comment) (models.Comment, error)
	GetAllComments() ([]models.Comment, error)
	GetCommentById(commentId int) (models.Comment, error)
	GetCommentsByGroupId(groupId int) (models.Comment, error)
	GetCommentsByUserId(userId int) (models.Comment, error)
	GetCommentsByPostId(postId int) (models.Comment, error)
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteAllComments() error
	DeleteCommentById(commentId int) error
	DeleteCommentsByGroupId(groupId int) error
	DeleteCommentsByUserId(userId int) error
	DeleteCommentsByPostId(postId int) error

	// Event
	CreateEvent(event models.Event) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
	GetEventById(eventId int) (models.Event, error)
	GetEventsByGroupId(groupId int) (models.Event, error)
	GetEventsByUserId(userId int) (models.Event, error)
	UpdateEvent(event models.Event) (models.Event, error)
	DeleteAllEvents() error
	DeleteEventById(eventId int) error
	DeleteEventsByGroupId(groupId int) error
	DeleteEventsByUserId(userId int) error

	// Message
	CreateMessage(message models.Message) (models.Message, error)
	GetAllMessages() ([]models.Message, error)
	GetMessagesByType(messageType string) ([]models.Message, error)
	GetMessageById(messageId int) (models.Message, error)
	GetMessagesBySenderId(senderId int) (models.Message, error)
	GetMessagesByTargetId(targetId int) (models.Message, error)
	UpdateMessage(message models.Message) (models.Message, error)
	DeleteAllMessages() error
	DeleteMessagesByType(messageType string) error
	DeleteMessageById(messageId int) error
	DeleteMessagesBySenderId(senderId int) error
	DeleteMessagesByTargetId(targetId int) error
}

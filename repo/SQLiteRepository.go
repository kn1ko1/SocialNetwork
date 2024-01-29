// Message from Matt
// Please do not touch this file

package repo

import (
	"errors"
	"log"
	"socialnetwork/models"
	"socialnetwork/sqlite"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	businessDb *sql.DB
	identityDb *sql.DB
}

func NewSQLiteRepository() *SQLiteRepository {
	ret := &SQLiteRepository{}
	db, err := sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ret.identityDb = db
	db, err = sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	ret.businessDb = db
	return ret
}

// Users
func (r *SQLiteRepository) CreateUser(user models.User) (models.User, error) {
	return sqlite.CreateUser(r.identityDb, user)
}
func (r *SQLiteRepository) GetAllUsers() ([]models.User, error) {
	return sqlite.GetAllUsers(r.identityDb)
}
func (r *SQLiteRepository) GetUserById(userId int) (models.User, error) {
	return sqlite.GetUserById(r.identityDb, userId)
}
func (r *SQLiteRepository) GetUserByEmail(email string) (models.User, error) {
	return sqlite.GetUserByEmail(r.identityDb, email)
}
func (r *SQLiteRepository) GetUserByUsername(username string) (models.User, error) {
	return sqlite.GetUserByUsername(r.identityDb, username)
}
func (r *SQLiteRepository) UpdateUser(user models.User) (models.User, error) {
	return sqlite.UpdateUser(r.identityDb, user)
}
func (r *SQLiteRepository) DeleteUserById(userId int) error {
	return sqlite.DeleteUserById(r.identityDb, userId)
}
func (r *SQLiteRepository) DeleteAllUsers() error {
	return sqlite.DeleteAllUsers(r.identityDb)
}

// Post
func (r *SQLiteRepository) CreatePost(post models.Post) (models.Post, error) {
	return sqlite.CreatePost(r.businessDb, post)
}
func (r *SQLiteRepository) GetAllPosts() ([]models.Post, error) {
	return sqlite.GetAllPosts(r.businessDb)
}
func (r *SQLiteRepository) GetPostById(postId int) (models.Post, error) {
	return sqlite.GetPostById(r.businessDb, postId)
}
func (r *SQLiteRepository) GetPostsByGroupId(groupId int) ([]models.Post, error) {
	return sqlite.GetPostsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	return sqlite.GetPostsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeletePostById(postId int) error {
	return sqlite.DeletePostById(r.businessDb, postId)
}
func (r *SQLiteRepository) UpdatePost(post models.Post) (models.Post, error) {
	return sqlite.UpdatePost(r.businessDb, post)
}
func (r *SQLiteRepository) DeletePostByGroupId(groupId int) error {
	return sqlite.DeletePostByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeletePostsByUserId(userId int) error {
	return sqlite.DeletePostsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteAllPosts() error {
	return sqlite.DeleteAllPosts(r.businessDb)
}

// Comments
func (r *SQLiteRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	return sqlite.CreateComment(r.businessDb, comment)
}
func (r *SQLiteRepository) GetAllComments() ([]models.Comment, error) {
	return sqlite.GetAllComments(r.businessDb)
}
func (r *SQLiteRepository) GetCommentById(commentId int) (models.Comment, error) {
	return sqlite.GetCommentById(r.businessDb, commentId)
}

//	func (r *SQLiteRepository) GetCommentsByGroupId(groupId int) ([]models.Comment, error) {
//		return sqlite.GetCommentsByGroupId(r.businessDb, groupId)
//	}
func (r *SQLiteRepository) GetCommentsByUserId(userId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	return sqlite.UpdateComment(r.businessDb, comment)
}
func (r *SQLiteRepository) DeleteCommentById(commentId int) error {
	return sqlite.DeleteCommentById(r.businessDb, commentId)
}
func (r *SQLiteRepository) DeleteCommentsByGroupId(groupId int) error {
	return sqlite.DeleteCommentsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteCommentsByUserId(userId int) error {
	return sqlite.DeleteCommentsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteCommentsByPostId(postId int) error {
	return sqlite.DeleteCommentsByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) DeleteAllComments() error {
	return sqlite.DeleteAllComments(r.businessDb)
}

// Event
func (r *SQLiteRepository) CreateEvent(event models.Event) (models.Event, error) {
	return sqlite.CreateEvent(r.businessDb, event)
}
func (r *SQLiteRepository) GetAllEvents() ([]models.Event, error) {
	return sqlite.GetAllEvents(r.businessDb)
}
func (r *SQLiteRepository) GetEventById(eventId int) (models.Event, error) {
	return sqlite.GetEventById(r.businessDb, eventId)
}
func (r *SQLiteRepository) GetEventsByGroupId(groupId int) ([]models.Event, error) {
	return sqlite.GetEventsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) GetEventsByUserId(userId int) ([]models.Event, error) {
	return sqlite.GetEventsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) UpdateEvent(event models.Event) (models.Event, error) {
	return sqlite.UpdateEvent(r.businessDb, event)
}
func (r *SQLiteRepository) DeleteEventById(eventId int) error {
	return sqlite.DeleteEventById(r.businessDb, eventId)
}
func (r *SQLiteRepository) DeleteEventsByGroupId(groupId int) error {
	return sqlite.DeleteEventsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteEventsByUserId(userId int) error {
	return sqlite.DeleteEventsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteAllEvents() error {
	return sqlite.DeleteAllEvents(r.businessDb)
}

// Message
func (r *SQLiteRepository) CreateMessage(message models.Message) (models.Message, error) {
	return sqlite.CreateMessage(r.businessDb, message)
}
func (r *SQLiteRepository) GetAllMessages() ([]models.Message, error) {
	return sqlite.GetAllMessages(r.businessDb)
}
func (r *SQLiteRepository) GetMessagesByType(messageType string) ([]models.Message, error) {
	return sqlite.GetMessagesByType(r.businessDb, messageType)
}
func (r *SQLiteRepository) GetMessageById(messageId int) (models.Message, error) {
	return sqlite.GetMessageById(r.businessDb, messageId)
}
func (r *SQLiteRepository) GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error) {
	return sqlite.GetMessagesBySenderAndTargetIDs(r.businessDb, senderId, targetId)
}
func (r *SQLiteRepository) UpdateMessage(message models.Message) (models.Message, error) {
	return sqlite.UpdateMessage(r.businessDb, message)
}
func (r *SQLiteRepository) DeleteMessagesByType(messageType string) error {
	return sqlite.DeleteMessagesByType(r.businessDb, messageType)
}
func (r *SQLiteRepository) DeleteMessageById(messageId int) error {
	return sqlite.DeleteMessageById(r.businessDb, messageId)
}
func (r *SQLiteRepository) DeleteMessagesBySenderId(senderId int) error {
	return sqlite.DeleteMessagesBySenderId(r.businessDb, senderId)
}
func (r *SQLiteRepository) DeleteMessagesByTargetId(targetId int) error {
	return sqlite.DeleteMessagesByTargetId(r.businessDb, targetId)
}
func (r *SQLiteRepository) DeleteAllMessages() error {
	return sqlite.DeleteAllMessages(r.businessDb)
}

// Group
func (r *SQLiteRepository) CreateGroup(group models.Group) (models.Group, error) {
	return group, errors.New("not implemented")
}

func (r *SQLiteRepository) GetAllGroups() ([]models.Group, error) {
	return nil, errors.New("not implemented")
}

func (r *SQLiteRepository) UpdateGroup(group models.Group) (models.Group, error) {
	return group, errors.New("not implemented")
}

func (r *SQLiteRepository) DeleteAllGroups() error {
	return errors.New("not implemented")
}

// Group_User
func (r *SQLiteRepository) CreateGroupUser(groupUser models.GroupUser) (models.GroupUser, error) {
	return groupUser, errors.New("not implemented")
}
func (r *SQLiteRepository) GetGroupUsersByUserId(userId int) ([]models.GroupUser, error) {
	var groupUsers []models.GroupUser
	return groupUsers, errors.New("not implemented")
}
func (r *SQLiteRepository) GetGroupUsersByGroupId(groupId int) ([]models.GroupUser, error) {
	return sqlite.GetGroupUsersByGroupId(r.businessDb, groupId)
}

// Notification
func (r *SQLiteRepository) CreateNotification(notification models.Notification) (models.Notification, error) {
	return notification, errors.New("not implemented")
}
func (r *SQLiteRepository) GetNotificationById(notificationId int) (models.Notification, error) {
	var emptyNotification models.Notification
	return emptyNotification, errors.New("not implemented")
}
func (r *SQLiteRepository) UpdateNotification(notification models.Notification) (models.Notification, error) {
	return notification, errors.New("not implemented")
}
func (r *SQLiteRepository) DeleteNotificationById(notificationId int) error {
	return errors.New("not implemented")
}

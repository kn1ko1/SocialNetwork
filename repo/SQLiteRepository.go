package repo

import (
	"database/sql"
	"socialnetwork/models"
	"socialnetwork/sqlite"
)

type DatabaseRepository struct {
}

func NewDatabaseRepository() *DatabaseRepository {
	return &DatabaseRepository{}
}

// Users
func (r *DatabaseRepository) CreateUser(database *sql.DB, user models.User) (models.User, error) {
	return sqlite.CreateUser(database, user)
}
func (r *DatabaseRepository) GetAllUsers(database *sql.DB) ([]models.User, error) {
	return sqlite.GetAllUsers(database)
}
func (r *DatabaseRepository) GetUserById(database *sql.DB, userId int) (models.User, error) {
	return sqlite.GetUserById(database, userId)
}
func (r *DatabaseRepository) GetUserByEmail(database *sql.DB, email string) (models.User, error) {
	return sqlite.GetUserByEmail(database, email)
}
func (r *DatabaseRepository) GetUserByUsername(database *sql.DB, username string) (models.User, error) {
	return sqlite.GetUserByUsername(database, username)
}
func (r *DatabaseRepository) UpdateUser(database *sql.DB, user models.User) (models.User, error) {
	return sqlite.UpdateUser(database, user)
}
func (r *DatabaseRepository) DeleteUserById(database *sql.DB, userId int) error {
	return sqlite.DeleteUserById(database, userId)
}
func (r *DatabaseRepository) DeleteAllUsers(database *sql.DB) error {
	return sqlite.DeleteAllUsers(database)
}

// Post
func (r *DatabaseRepository) CreatePost(database *sql.DB, post models.Post) (models.Post, error) {
	return sqlite.CreatePost(database, post)
}
func (r *DatabaseRepository) GetAllPosts(database *sql.DB) ([]models.Post, error) {
	return sqlite.GetAllPosts(database)
}
func (r *DatabaseRepository) GetPostById(database *sql.DB, postId int) (models.Post, error) {
	return sqlite.GetPostById(database, postId)
}
func (r *DatabaseRepository) GetPostsByGroupId(database *sql.DB, groupId int) ([]models.Post, error) {
	return sqlite.GetPostsByGroupId(database, groupId)
}
func (r *DatabaseRepository) GetPostsByUserId(database *sql.DB, userId int) ([]models.Post, error) {
	return sqlite.GetPostsByUserId(database, userId)
}
func (r *DatabaseRepository) DeletePostById(database *sql.DB, postId int) error {
	return sqlite.DeletePostById(database, postId)
}
func (r *DatabaseRepository) UpdatePost(database *sql.DB, post models.Post) (models.Post, error) {
	return sqlite.UpdatePost(database, post)
}
func (r *DatabaseRepository) DeletePostByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeletePostByGroupId(database, groupId)
}
func (r *DatabaseRepository) DeletePostsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeletePostsByUserId(database, userId)
}
func (r *DatabaseRepository) DeleteAllPosts(database *sql.DB) error {
	return sqlite.DeleteAllPosts(database)
}

// Comments
func (r *DatabaseRepository) CreateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {
	return sqlite.CreateComment(database, comment)
}
func (r *DatabaseRepository) GetAllComments(database *sql.DB) ([]models.Comment, error) {
	return sqlite.GetAllComments(database)
}
func (r *DatabaseRepository) GetCommentById(database *sql.DB, commentId int) (models.Comment, error) {
	return sqlite.GetCommentById(database, commentId)
}
func (r *DatabaseRepository) GetCommentsByGroupId(database *sql.DB, groupId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByGroupId(database, groupId)
}
func (r *DatabaseRepository) GetCommentsByUserId(database *sql.DB, userId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByUserId(database, userId)
}
func (r *DatabaseRepository) GetCommentsByPostId(database *sql.DB, postId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByPostId(database, postId)
}
func (r *DatabaseRepository) UpdateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {
	return sqlite.UpdateComment(database, comment)
}
func (r *DatabaseRepository) DeleteCommentById(database *sql.DB, commentId int) error {
	return sqlite.DeleteCommentById(database, commentId)
}
func (r *DatabaseRepository) DeleteCommentsByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeleteCommentsByGroupId(database, groupId)
}
func (r *DatabaseRepository) DeleteCommentsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeleteCommentsByUserId(database, userId)
}
func (r *DatabaseRepository) DeleteCommentsByPostId(database *sql.DB, postId int) error {
	return sqlite.DeleteCommentsByPostId(database, postId)
}
func (r *DatabaseRepository) DeleteAllComments(database *sql.DB) error {
	return sqlite.DeleteAllComments(database)
}

// Event
func (r *DatabaseRepository) CreateEvent(database *sql.DB, event models.Event) (models.Event, error) {
	return sqlite.CreateEvent(database, event)
}
func (r *DatabaseRepository) GetAllEvents(database *sql.DB) ([]models.Event, error) {
	return sqlite.GetAllEvents(database)
}
func (r *DatabaseRepository) GetEventById(database *sql.DB, eventId int) (models.Event, error) {
	return sqlite.GetEventById(database, eventId)
}
func (r *DatabaseRepository) GetEventsByGroupId(database *sql.DB, groupId int) ([]models.Event, error) {
	return sqlite.GetEventsByGroupId(database, groupId)
}
func (r *DatabaseRepository) GetEventsByUserId(database *sql.DB, userId int) ([]models.Event, error) {
	return sqlite.GetEventsByUserId(database, userId)
}
func (r *DatabaseRepository) UpdateEvent(database *sql.DB, event models.Event) (models.Event, error) {
	return sqlite.UpdateEvent(database, event)
}
func (r *DatabaseRepository) DeleteEventById(database *sql.DB, eventId int) error {
	return sqlite.DeleteEventById(database, eventId)
}
func (r *DatabaseRepository) DeleteEventsByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeleteEventsByGroupId(database, groupId)
}
func (r *DatabaseRepository) DeleteEventsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeleteEventsByUserId(database, userId)
}
func (r *DatabaseRepository) DeleteAllEvents(database *sql.DB) error {
	return sqlite.DeleteAllEvents(database)
}

// Message
func (r *DatabaseRepository) CreateMessage(database *sql.DB, message models.Message) (models.Message, error) {
	return sqlite.CreateMessage(database, message)
}
func (r *DatabaseRepository) GetAllMessages(database *sql.DB) ([]models.Message, error) {
	return sqlite.GetAllMessages(database)
}
func (r *DatabaseRepository) GetMessagesByType(database *sql.DB, messageType string) ([]models.Message, error) {
	return sqlite.GetMessagesByType(database, messageType)
}
func (r *DatabaseRepository) GetMessageById(database *sql.DB, messageId int) (models.Message, error) {
	return sqlite.GetMessageById(database, messageId)
}
func (r *DatabaseRepository) GetMessagesBySenderAndTargetIDs(database *sql.DB, senderId, targetId int) ([]models.Message, error) {
	return sqlite.GetMessagesBySenderAndTargetIDs(database, senderId, targetId)
}
func (r *DatabaseRepository) UpdateMessage(database *sql.DB, message models.Message) (models.Message, error) {
	return sqlite.UpdateMessage(database, message)
}
func (r *DatabaseRepository) DeleteMessagesByType(database *sql.DB, messageType string) error {
	return sqlite.DeleteMessagesByType(database, messageType)
}
func (r *DatabaseRepository) DeleteMessageById(database *sql.DB, messageId int) error {
	return sqlite.DeleteMessageById(database, messageId)
}
func (r *DatabaseRepository) DeleteMessagesBySenderId(database *sql.DB, senderId int) error {
	return sqlite.DeleteMessagesBySenderId(database, senderId)
}
func (r *DatabaseRepository) DeleteMessagesByTargetId(database *sql.DB, targetId int) error {
	return sqlite.DeleteMessagesByTargetId(database, targetId)
}
func (r *DatabaseRepository) DeleteAllMessages(database *sql.DB) error {
	return sqlite.DeleteAllMessages(database)
}

package repo

import (
	"database/sql"
	"log"
	"socialnetwork/models"
	"socialnetwork/sqlite"
)

type SQLiteRepository struct {
	businessDb *sql.DB
	identityDb *sql.DB
}

func NewDatabaseRepository() *SQLiteRepository {
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
func (r *SQLiteRepository) GetAllUsers(database *sql.DB) ([]models.User, error) {
	return sqlite.GetAllUsers(database)
}
func (r *SQLiteRepository) GetUserById(database *sql.DB, userId int) (models.User, error) {
	return sqlite.GetUserById(database, userId)
}
func (r *SQLiteRepository) GetUserByEmail(database *sql.DB, email string) (models.User, error) {
	return sqlite.GetUserByEmail(database, email)
}
func (r *SQLiteRepository) GetUserByUsername(database *sql.DB, username string) (models.User, error) {
	return sqlite.GetUserByUsername(database, username)
}
func (r *SQLiteRepository) UpdateUser(database *sql.DB, user models.User) (models.User, error) {
	return sqlite.UpdateUser(database, user)
}
func (r *SQLiteRepository) DeleteUserById(database *sql.DB, userId int) error {
	return sqlite.DeleteUserById(database, userId)
}
func (r *SQLiteRepository) DeleteAllUsers(database *sql.DB) error {
	return sqlite.DeleteAllUsers(database)
}

// Post
func (r *SQLiteRepository) CreatePost(database *sql.DB, post models.Post) (models.Post, error) {
	return sqlite.CreatePost(database, post)
}
func (r *SQLiteRepository) GetAllPosts(database *sql.DB) ([]models.Post, error) {
	return sqlite.GetAllPosts(database)
}
func (r *SQLiteRepository) GetPostById(database *sql.DB, postId int) (models.Post, error) {
	return sqlite.GetPostById(database, postId)
}
func (r *SQLiteRepository) GetPostsByGroupId(database *sql.DB, groupId int) ([]models.Post, error) {
	return sqlite.GetPostsByGroupId(database, groupId)
}
func (r *SQLiteRepository) GetPostsByUserId(database *sql.DB, userId int) ([]models.Post, error) {
	return sqlite.GetPostsByUserId(database, userId)
}
func (r *SQLiteRepository) DeletePostById(database *sql.DB, postId int) error {
	return sqlite.DeletePostById(database, postId)
}
func (r *SQLiteRepository) UpdatePost(database *sql.DB, post models.Post) (models.Post, error) {
	return sqlite.UpdatePost(database, post)
}
func (r *SQLiteRepository) DeletePostByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeletePostByGroupId(database, groupId)
}
func (r *SQLiteRepository) DeletePostsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeletePostsByUserId(database, userId)
}
func (r *SQLiteRepository) DeleteAllPosts(database *sql.DB) error {
	return sqlite.DeleteAllPosts(database)
}

// Comments
func (r *SQLiteRepository) CreateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {
	return sqlite.CreateComment(database, comment)
}
func (r *SQLiteRepository) GetAllComments(database *sql.DB) ([]models.Comment, error) {
	return sqlite.GetAllComments(database)
}
func (r *SQLiteRepository) GetCommentById(database *sql.DB, commentId int) (models.Comment, error) {
	return sqlite.GetCommentById(database, commentId)
}
func (r *SQLiteRepository) GetCommentsByGroupId(database *sql.DB, groupId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByGroupId(database, groupId)
}
func (r *SQLiteRepository) GetCommentsByUserId(database *sql.DB, userId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByUserId(database, userId)
}
func (r *SQLiteRepository) GetCommentsByPostId(database *sql.DB, postId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByPostId(database, postId)
}
func (r *SQLiteRepository) UpdateComment(database *sql.DB, comment models.Comment) (models.Comment, error) {
	return sqlite.UpdateComment(database, comment)
}
func (r *SQLiteRepository) DeleteCommentById(database *sql.DB, commentId int) error {
	return sqlite.DeleteCommentById(database, commentId)
}
func (r *SQLiteRepository) DeleteCommentsByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeleteCommentsByGroupId(database, groupId)
}
func (r *SQLiteRepository) DeleteCommentsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeleteCommentsByUserId(database, userId)
}
func (r *SQLiteRepository) DeleteCommentsByPostId(database *sql.DB, postId int) error {
	return sqlite.DeleteCommentsByPostId(database, postId)
}
func (r *SQLiteRepository) DeleteAllComments(database *sql.DB) error {
	return sqlite.DeleteAllComments(database)
}

// Event
func (r *SQLiteRepository) CreateEvent(database *sql.DB, event models.Event) (models.Event, error) {
	return sqlite.CreateEvent(database, event)
}
func (r *SQLiteRepository) GetAllEvents(database *sql.DB) ([]models.Event, error) {
	return sqlite.GetAllEvents(database)
}
func (r *SQLiteRepository) GetEventById(database *sql.DB, eventId int) (models.Event, error) {
	return sqlite.GetEventById(database, eventId)
}
func (r *SQLiteRepository) GetEventsByGroupId(database *sql.DB, groupId int) ([]models.Event, error) {
	return sqlite.GetEventsByGroupId(database, groupId)
}
func (r *SQLiteRepository) GetEventsByUserId(database *sql.DB, userId int) ([]models.Event, error) {
	return sqlite.GetEventsByUserId(database, userId)
}
func (r *SQLiteRepository) UpdateEvent(database *sql.DB, event models.Event) (models.Event, error) {
	return sqlite.UpdateEvent(database, event)
}
func (r *SQLiteRepository) DeleteEventById(database *sql.DB, eventId int) error {
	return sqlite.DeleteEventById(database, eventId)
}
func (r *SQLiteRepository) DeleteEventsByGroupId(database *sql.DB, groupId int) error {
	return sqlite.DeleteEventsByGroupId(database, groupId)
}
func (r *SQLiteRepository) DeleteEventsByUserId(database *sql.DB, userId int) error {
	return sqlite.DeleteEventsByUserId(database, userId)
}
func (r *SQLiteRepository) DeleteAllEvents(database *sql.DB) error {
	return sqlite.DeleteAllEvents(database)
}

// Message
func (r *SQLiteRepository) CreateMessage(database *sql.DB, message models.Message) (models.Message, error) {
	return sqlite.CreateMessage(database, message)
}
func (r *SQLiteRepository) GetAllMessages(database *sql.DB) ([]models.Message, error) {
	return sqlite.GetAllMessages(database)
}
func (r *SQLiteRepository) GetMessagesByType(database *sql.DB, messageType string) ([]models.Message, error) {
	return sqlite.GetMessagesByType(database, messageType)
}
func (r *SQLiteRepository) GetMessageById(database *sql.DB, messageId int) (models.Message, error) {
	return sqlite.GetMessageById(database, messageId)
}
func (r *SQLiteRepository) GetMessagesBySenderAndTargetIDs(database *sql.DB, senderId, targetId int) ([]models.Message, error) {
	return sqlite.GetMessagesBySenderAndTargetIDs(database, senderId, targetId)
}
func (r *SQLiteRepository) UpdateMessage(database *sql.DB, message models.Message) (models.Message, error) {
	return sqlite.UpdateMessage(database, message)
}
func (r *SQLiteRepository) DeleteMessagesByType(database *sql.DB, messageType string) error {
	return sqlite.DeleteMessagesByType(database, messageType)
}
func (r *SQLiteRepository) DeleteMessageById(database *sql.DB, messageId int) error {
	return sqlite.DeleteMessageById(database, messageId)
}
func (r *SQLiteRepository) DeleteMessagesBySenderId(database *sql.DB, senderId int) error {
	return sqlite.DeleteMessagesBySenderId(database, senderId)
}
func (r *SQLiteRepository) DeleteMessagesByTargetId(database *sql.DB, targetId int) error {
	return sqlite.DeleteMessagesByTargetId(database, targetId)
}
func (r *SQLiteRepository) DeleteAllMessages(database *sql.DB) error {
	return sqlite.DeleteAllMessages(database)
}

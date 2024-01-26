package repo

import (
	"database/sql"
	"errors"
	"log"
	"socialnetwork/models"
	"socialnetwork/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sutTableRuns = 10
)

var (
	validUser = *models.GenerateValidUser()
	validPost = *models.GenerateValidPost()
)

// Field commented for implementation example reasons:
// Should take as a field a private DB instance
//
// The repo instance is then responsible for maintaining
// the concrete DB connection
type DummyRepository struct {
	identityDb *sql.DB
	businessDb *sql.DB
}

// Constructor function
func NewDummyRepository() *DummyRepository {
	// The DB field would be constructed properly here!
	identityDb, err := sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatal(err)
	}
	businessDb, err := sql.Open(dbDriver, businessDbPath)
	if err != nil {
		log.Fatal(err)
	}
	return &DummyRepository{identityDb: identityDb, businessDb: businessDb}
}

// Sim user creation
func (r *DummyRepository) CreateUser(user models.User) (models.User, error) {
	user.UserId = 1
	return user, nil
}

// Sim User retrieval
func (r *DummyRepository) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		users[i] = *models.GenerateValidUser()
	}
	return users, nil
}
func (r *DummyRepository) GetUserById(userId int) (models.User, error) {
	user := validUser
	user.UserId = userId
	return user, nil
}

func (r *DummyRepository) GetUserByEmail(email string) (models.User, error) {
	user := validUser
	user.Email = email
	return user, nil
}

func (r *DummyRepository) GetUserByUsername(username string) (models.User, error) {
	user := validUser
	user.Username = username
	return user, nil
}

func (r *DummyRepository) UpdateUser(user models.User) (models.User, error) {
	return user, nil
}

func (r *DummyRepository) DeleteUserById(userId int) error {
	return sqlite.DeleteUserById(r.identityDb, userId)
}

func (r *DummyRepository) DeleteAllUsers() error {
	return sqlite.DeleteAllUsers(r.identityDb)
}

// Post
func (r *DummyRepository) CreatePost(post models.Post) (models.Post, error) {
	return sqlite.CreatePost(r.businessDb, post)
}
func (r *DummyRepository) GetAllPosts() ([]models.Post, error) {
	return sqlite.GetAllPosts(r.businessDb)
}
func (r *DummyRepository) GetPostById(postId int) (models.Post, error) {
	return sqlite.GetPostById(r.businessDb, postId)
}
func (r *DummyRepository) GetPostsByGroupId(groupId int) ([]models.Post, error) {
	return sqlite.GetPostsByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	return sqlite.GetPostsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) DeletePostById(postId int) error {
	return sqlite.DeletePostById(r.businessDb, postId)
}
func (r *DummyRepository) UpdatePost(post models.Post) (models.Post, error) {
	return sqlite.UpdatePost(r.businessDb, post)
}
func (r *DummyRepository) DeletePostByGroupId(groupId int) error {
	return sqlite.DeletePostByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) DeletePostsByUserId(userId int) error {
	return sqlite.DeletePostsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) DeleteAllPosts() error {
	return sqlite.DeleteAllPosts(r.businessDb)
}

// Comments
func (r *DummyRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	return sqlite.CreateComment(r.businessDb, comment)
}
func (r *DummyRepository) GetAllComments() ([]models.Comment, error) {
	return sqlite.GetAllComments(r.businessDb)
}
func (r *DummyRepository) GetCommentById(commentId int) (models.Comment, error) {
	return sqlite.GetCommentById(r.businessDb, commentId)
}
func (r *DummyRepository) GetCommentsByGroupId(groupId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) GetCommentsByUserId(userId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	return sqlite.GetCommentsByPostId(r.businessDb, postId)
}
func (r *DummyRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	return sqlite.UpdateComment(r.businessDb, comment)
}
func (r *DummyRepository) DeleteCommentById(commentId int) error {
	return sqlite.DeleteCommentById(r.businessDb, commentId)
}
func (r *DummyRepository) DeleteCommentsByGroupId(groupId int) error {
	return sqlite.DeleteCommentsByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) DeleteCommentsByUserId(userId int) error {
	return sqlite.DeleteCommentsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) DeleteCommentsByPostId(postId int) error {
	return sqlite.DeleteCommentsByPostId(r.businessDb, postId)
}
func (r *DummyRepository) DeleteAllComments() error {
	return sqlite.DeleteAllComments(r.businessDb)
}

// Event
func (r *DummyRepository) CreateEvent(event models.Event) (models.Event, error) {
	return sqlite.CreateEvent(r.businessDb, event)
}
func (r *DummyRepository) GetAllEvents() ([]models.Event, error) {
	return sqlite.GetAllEvents(r.businessDb)
}
func (r *DummyRepository) GetEventById(eventId int) (models.Event, error) {
	return sqlite.GetEventById(r.businessDb, eventId)
}
func (r *DummyRepository) GetEventsByGroupId(groupId int) ([]models.Event, error) {
	return sqlite.GetEventsByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) GetEventsByUserId(userId int) ([]models.Event, error) {
	return sqlite.GetEventsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) UpdateEvent(event models.Event) (models.Event, error) {
	return sqlite.UpdateEvent(r.businessDb, event)
}
func (r *DummyRepository) DeleteEventById(eventId int) error {
	return sqlite.DeleteEventById(r.businessDb, eventId)
}
func (r *DummyRepository) DeleteEventsByGroupId(groupId int) error {
	return sqlite.DeleteEventsByGroupId(r.businessDb, groupId)
}
func (r *DummyRepository) DeleteEventsByUserId(userId int) error {
	return sqlite.DeleteEventsByUserId(r.businessDb, userId)
}
func (r *DummyRepository) DeleteAllEvents() error {
	return sqlite.DeleteAllEvents(r.businessDb)
}

// Message
func (r *DummyRepository) CreateMessage(message models.Message) (models.Message, error) {
	return sqlite.CreateMessage(r.businessDb, message)
}
func (r *DummyRepository) GetAllMessages() ([]models.Message, error) {
	return sqlite.GetAllMessages(r.businessDb)
}
func (r *DummyRepository) GetMessagesByType(messageType string) ([]models.Message, error) {
	return sqlite.GetMessagesByType(r.businessDb, messageType)
}
func (r *DummyRepository) GetMessageById(messageId int) (models.Message, error) {
	return sqlite.GetMessageById(r.businessDb, messageId)
}
func (r *DummyRepository) GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error) {
	return sqlite.GetMessagesBySenderAndTargetIDs(r.businessDb, senderId, targetId)
}
func (r *DummyRepository) UpdateMessage(message models.Message) (models.Message, error) {
	return sqlite.UpdateMessage(r.businessDb, message)
}
func (r *DummyRepository) DeleteMessagesByType(messageType string) error {
	return nil
}
func (r *DummyRepository) DeleteMessageById(messageId int) error {
	return nil
}
func (r *DummyRepository) DeleteMessagesBySenderId(senderId int) error {
	return nil
}
func (r *DummyRepository) DeleteMessagesByTargetId(targetId int) error {
	return nil
}
func (r *DummyRepository) DeleteAllMessages() error {
	return nil
}

func (r *DummyRepository) CreateGroup(group models.Group) (models.Group, error) {
	return group, errors.New("not implemented")
}

func (r *DummyRepository) GetAllGroups() ([]models.Group, error) {
	return nil, errors.New("not implemented")
}

func (r *DummyRepository) UpdateGroup(group models.Group) (models.Group, error) {
	return group, errors.New("not implemented")
}

func (r *DummyRepository) DeleteAllGroups() error {
	return errors.New("not implemented")
}

package repo

import (
	"database/sql"
	"log"
	"socialnetwork/models"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sutTableRuns = 10
)

var (
	validUser    = *models.GenerateValidUser()
	validPost    = *models.GenerateValidPost()
	validComment = *models.GenerateValidComment()
	validEvent   = *models.GenerateValidEvent()
	validMessage = *models.GenerateValidMessage()
	validGroup   = *models.GenerateValidGroup()
)

func init() {
	validUser.UserId = 1
	validPost.PostId = 1
	validComment.CommentId = 1
	validEvent.EventId = 1
	validMessage.MessageId = 1
	validGroup.GroupID = 1
}

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
		u := *models.GenerateValidUser()
		u.UserId = i + 1
		users[i] = u
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
	return nil
}

func (r *DummyRepository) DeleteAllUsers() error {
	return nil
}

// Post
func (r *DummyRepository) CreatePost(post models.Post) (models.Post, error) {
	post.PostId = 1
	return post, nil
}

func (r *DummyRepository) GetAllPosts() ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := *models.GenerateValidPost()
		p.PostId = i + 1
		posts[i] = p
	}
	return posts, nil
}

func (r *DummyRepository) GetPostById(postId int) (models.Post, error) {
	post := validPost
	post.PostId = postId
	return post, nil
}

func (r *DummyRepository) GetPostsByGroupId(groupId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := *models.GenerateValidPost()
		p.PostId = i + 1
		p.GroupId = groupId
		posts[i] = p
	}
	return posts, nil
}

func (r *DummyRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := *models.GenerateValidPost()
		p.PostId = i + 1
		p.UserId = userId
		posts[i] = p
	}
	return posts, nil
}
func (r *DummyRepository) DeletePostById(postId int) error {
	return nil
}
func (r *DummyRepository) UpdatePost(post models.Post) (models.Post, error) {
	return post, nil
}
func (r *DummyRepository) DeletePostByGroupId(groupId int) error {
	return nil
}
func (r *DummyRepository) DeletePostsByUserId(userId int) error {
	return nil
}
func (r *DummyRepository) DeleteAllPosts() error {
	return nil
}

// Comments
func (r *DummyRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	comment.CommentId = 1
	return comment, nil
}
func (r *DummyRepository) GetAllComments() ([]models.Comment, error) {
	comments := make([]models.Comment, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		c := *models.GenerateValidComment()
		c.CommentId = i + 1
		comments[i] = c
	}
	return comments, nil
}

func (r *DummyRepository) GetCommentById(commentId int) (models.Comment, error) {
	comment := validComment
	comment.CommentId = commentId
	return comment, nil
}

func (r *DummyRepository) GetCommentsByUserId(userId int) ([]models.Comment, error) {
	comments := make([]models.Comment, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		c := *models.GenerateValidComment()
		c.CommentId = i + 1
		c.UserId = userId
		comments[i] = c
	}
	return comments, nil
}

func (r *DummyRepository) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	comments := make([]models.Comment, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		c := *models.GenerateValidComment()
		c.CommentId = i + 1
		c.PostId = postId
		comments[i] = c
	}
	return comments, nil
}

func (r *DummyRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	return comment, nil
}

func (r *DummyRepository) DeleteCommentById(commentId int) error {
	return nil
}

func (r *DummyRepository) DeleteCommentsByGroupId(groupId int) error {
	return nil
}

func (r *DummyRepository) DeleteCommentsByUserId(userId int) error {
	return nil
}

func (r *DummyRepository) DeleteCommentsByPostId(postId int) error {
	return nil
}

func (r *DummyRepository) DeleteAllComments() error {
	return nil
}

// Event
func (r *DummyRepository) CreateEvent(event models.Event) (models.Event, error) {
	event.EventId = 1
	return event, nil
}
func (r *DummyRepository) GetAllEvents() ([]models.Event, error) {
	events := make([]models.Event, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		e := *models.GenerateValidEvent()
		e.EventId = i + 1
		events[i] = e
	}
	return events, nil
}

func (r *DummyRepository) GetEventById(eventId int) (models.Event, error) {
	event := validEvent
	event.EventId = eventId
	return event, nil
}
func (r *DummyRepository) GetEventsByGroupId(groupId int) ([]models.Event, error) {
	events := make([]models.Event, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		e := *models.GenerateValidEvent()
		e.EventId = i + 1
		e.GroupId = groupId
		events[i] = e
	}
	return events, nil
}
func (r *DummyRepository) GetEventsByUserId(userId int) ([]models.Event, error) {
	events := make([]models.Event, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		e := *models.GenerateValidEvent()
		e.EventId = i + 1
		e.UserId = userId
		events[i] = e
	}
	return events, nil
}
func (r *DummyRepository) UpdateEvent(event models.Event) (models.Event, error) {
	return event, nil
}
func (r *DummyRepository) DeleteEventById(eventId int) error {
	return nil
}
func (r *DummyRepository) DeleteEventsByGroupId(groupId int) error {
	return nil
}
func (r *DummyRepository) DeleteEventsByUserId(userId int) error {
	return nil
}
func (r *DummyRepository) DeleteAllEvents() error {
	return nil
}

// Message
func (r *DummyRepository) CreateMessage(message models.Message) (models.Message, error) {
	message.MessageId = 1
	return message, nil
}
func (r *DummyRepository) GetAllMessages() ([]models.Message, error) {
	messages := make([]models.Message, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		m := *models.GenerateValidMessage()
		m.MessageId = i + 1
		messages[i] = m
	}
	return messages, nil
}
func (r *DummyRepository) GetMessagesByType(messageType string) ([]models.Message, error) {
	messages := make([]models.Message, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		m := *models.GenerateValidMessage()
		m.MessageId = i + 1
		m.MessageType = messageType
		messages[i] = m
	}
	return messages, nil
}
func (r *DummyRepository) GetMessageById(messageId int) (models.Message, error) {
	message := validMessage
	message.MessageId = messageId
	return message, nil
}
func (r *DummyRepository) GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error) {
	messages := make([]models.Message, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		m := *models.GenerateValidMessage()
		m.MessageId = i + 1
		m.SenderId = senderId
		m.TargetId = targetId
		messages[i] = m
	}
	return messages, nil
}
func (r *DummyRepository) UpdateMessage(message models.Message) (models.Message, error) {
	return message, nil
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
	group.GroupID = 1
	return group, nil
}

func (r *DummyRepository) GetAllGroups() ([]models.Group, error) {
	groups := make([]models.Group, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		g := *models.GenerateValidGroup()
		g.GroupID = i + 1
		groups[i] = g
	}
	return groups, nil
}

func (r *DummyRepository) UpdateGroup(group models.Group) (models.Group, error) {
	return group, nil
}

func (r *DummyRepository) DeleteAllGroups() error {
	return nil
}

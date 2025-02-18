package repo

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"socialnetwork/Server/models"
	"socialnetwork/Server/transport"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sutTableRuns = 10
)

var (
	validUser         = *models.GenerateValidUser()
	validUserUser     = *models.GenerateValidUserUser()
	validPost         = *models.GenerateValidPost()
	validComment      = *models.GenerateValidComment()
	validEvent        = *models.GenerateValidEvent()
	validMessage      = *models.GenerateValidMessage()
	validGroup        = *models.GenerateValidGroup()
	validGroupUser    = *models.GenerateValidGroupUser()
	validNotification = *models.GenerateValidNotification()
)

func init() {
	validUser.UserId = rand.Intn(101)
	validPost.PostId = rand.Intn(101)
	validComment.CommentId = rand.Intn(101)
	validEvent.EventId = rand.Intn(101)
	validMessage.MessageId = rand.Intn(101)
	validGroup.GroupId = rand.Intn(101)
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

func (r *DummyRepository) GetProfileDataForUser(userId int) (transport.ProfileModel, error) {
	var profileModel transport.ProfileModel
	return profileModel, errors.New("not implimented yet")
}

func (r *DummyRepository) UpdateIsPublic(userId int, isPublic bool) error {
	return errors.New("not implimented yet")
}

// Users

func (r *DummyRepository) CreateUser(user models.User) (models.User, error) {
	user.UserId = 1
	return user, nil
}
func (r *DummyRepository) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		u := validUser
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
func (r *DummyRepository) GetUsersByPublic() ([]models.User, error) {
	users := make([]models.User, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		u := validUser
		u.UserId = i + 1
		u.IsPublic = true
		users[i] = u
	}
	return users, nil
}

//	func (r *DummyRepository) GetUserByEmail(email string) (models.User, error) {
//		user := validUser
//		user.Email = email
//		return user, nil
//	}
//
//	func (r *DummyRepository) GetUserByUsername(username string) (models.User, error) {
//		user := validUser
//		user.Username = username
//		return user, nil
//	}
func (r *DummyRepository) GetUserByUsernameOrEmail(usernameOrEmail string) (models.User, error) {
	return *models.GenerateValidUser(), nil
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

// UserUser
func (r *DummyRepository) CreateUserUser(userUser models.UserUser) (models.UserUser, error) {
	return validUserUser, nil
}
func (r *DummyRepository) GetUserUsersBySubjectId(subjectId int) ([]models.UserUser, error) {
	var userUsers []models.UserUser
	for i := 0; i < 5; i++ {
		user := validUserUser
		user.SubjectId = subjectId
		userUsers = append(userUsers, user)
	}

	return userUsers, nil
}
func (r *DummyRepository) GetUserUsersByFollowerId(followerId int) ([]models.UserUser, error) {
	var userUsers []models.UserUser
	for i := 0; i < 5; i++ {
		user := validUserUser
		user.FollowerId = followerId
		userUsers = append(userUsers, user)
	}

	return userUsers, nil
}

func (r *DummyRepository) GetUserUserByFollowerIdAndSubjectId(followerId, subjectId int) (models.UserUser, error) {

	userUser := validUserUser
	userUser.FollowerId = followerId
	userUser.SubjectId = subjectId

	return userUser, nil
}

func (r *DummyRepository) DeleteUserUsersBySubjectId(subjectId int) error {
	return nil
}
func (r *DummyRepository) DeleteUserUsersByFollowerId(followerId int) error {
	return nil
}
func (r *DummyRepository) DeleteUserUserBySubjectIdAndFollowerId(subjectId, followerId int) error {
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
		p := validPost
		p.PostId = i + 1
		posts[i] = p
	}
	return posts, nil
}
func (r *DummyRepository) GetPostById(postId int) (models.Post, error) {
	post := validPost
	post.PostId = postId
	err := post.Validate()
	log.Println("error in dummy GetPostById.", err)
	if err != nil {
		var emptypost models.Post
		return emptypost, err
	}
	return post, nil
}
func (r *DummyRepository) GetPostsByGroupId(groupId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := validPost
		p.PostId = i + 1
		p.GroupId = groupId
		posts[i] = p
	}
	return posts, nil
}
func (r *DummyRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := validPost
		p.PostId = i + 1
		p.UserId = userId
		posts[i] = p
		err := p.Validate()
		log.Println("error in dummy GetPostById.", err)
		if err != nil {
			var emptyposts []models.Post
			return emptyposts, err
		}
	}
	return posts, nil
}
func (r *DummyRepository) GetPostsAlmostPrivateForUserId(userId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := validPost
		p.UserId = userId
		p.PostId = i + 1
		posts[i] = p
	}
	return posts, nil
}
func (r *DummyRepository) GetPostsPrivateForUserId(userId int) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := validPost
		p.UserId = userId
		p.PostId = i + 1
		posts[i] = p
	}
	return posts, nil
}

func (r *DummyRepository) GetPostsByPrivacy(privacy string) ([]models.Post, error) {
	posts := make([]models.Post, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		p := validPost
		p.Privacy = privacy
		p.PostId = i + 1
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

// Post_Users
func (r *DummyRepository) CreatePostUser(postUser models.PostUser) (models.PostUser, error) {

	return postUser, nil
}
func (r *DummyRepository) GetPostUsersByUserId(userId int) ([]models.PostUser, error) {
	var postUsers []models.PostUser
	return postUsers, errors.New("not implimented yet")
}

func (r *DummyRepository) GetPostUsersByPostId(postId int) ([]models.PostUser, error) {
	var postUsers []models.PostUser
	return postUsers, errors.New("not implimented yet")
}

func (r *DummyRepository) DeletePostUsersByUserId(userId int) error {
	return nil
}

func (r *DummyRepository) DeletePostUsersByPostId(postId int) error {
	return nil
}

func (r *DummyRepository) DeletePostUserByPostIdAndUserId(postId, userId int) error {
	return nil
}

func (r *DummyRepository) DeleteAllPostUsers() error {
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
		c := validComment
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
		c := validComment
		c.CommentId = i + 1
		c.UserId = userId
		comments[i] = c
	}
	return comments, nil
}
func (r *DummyRepository) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	comments := make([]models.Comment, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		c := validComment
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
		e := validEvent
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
		e := validEvent
		e.EventId = i + 1
		e.GroupId = groupId
		events[i] = e
	}
	return events, nil
}
func (r *DummyRepository) GetEventsByUserId(userId int) ([]models.Event, error) {
	events := make([]models.Event, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		e := validEvent
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

// EventUsers
func (r *DummyRepository) CreateEventUser(eventUser models.EventUser) (models.EventUser, error) {
	return eventUser, nil
}
func (r *DummyRepository) GetEventUsersByUserId(userId int) ([]models.EventUser, error) {
	ctime := time.Now().UTC().UnixMilli()
	eventUsers := []models.EventUser{}
	user1 := models.EventUser{
		EventUserId: 1,
		CreatedAt:   ctime,
		EventId:     1,
		UpdatedAt:   ctime,
		UserId:      userId,
	}

	user2 := models.EventUser{
		EventUserId: 2,
		CreatedAt:   ctime,
		EventId:     2,
		UpdatedAt:   ctime,
		UserId:      userId,
	}
	eventUsers = append(eventUsers, user1, user2)

	return eventUsers, nil
}
func (r *DummyRepository) GetEventUsersByEventId(eventId int) ([]models.EventUser, error) {
	ctime := time.Now().UTC().UnixMilli()
	eventUsers := []models.EventUser{}
	user1 := models.EventUser{
		EventUserId: 1,
		CreatedAt:   ctime,
		EventId:     eventId,
		UpdatedAt:   ctime,
		UserId:      1,
	}

	user2 := models.EventUser{
		EventUserId: 2,
		CreatedAt:   ctime,
		EventId:     eventId,
		UpdatedAt:   ctime,
		UserId:      3,
	}
	eventUsers = append(eventUsers, user1, user2)

	return eventUsers, nil
}
func (r *DummyRepository) DeleteEventUsersByUserId(userId int) error {
	return nil
}
func (r *DummyRepository) DeleteEventUsersByEventId(eventId int) error {
	return nil
}
func (r *DummyRepository) DeleteEventUserByEventIdAndUserId(eventId, userId int) error {
	return nil
}
func (r *DummyRepository) DeleteAllEventUsers() error {
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
		m := validMessage
		m.MessageId = i + 1
		messages[i] = m
	}
	return messages, nil
}
func (r *DummyRepository) GetMessagesByMessageTypeandTargetId(messageType string, targetId int) ([]models.Message, error) {
	messages := make([]models.Message, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		m := validMessage
		m.MessageId = i + 1
		m.MessageType = messageType
		m.TargetId = targetId
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
		m := validMessage
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

// Group
func (r *DummyRepository) CreateGroup(group models.Group) (models.Group, error) {
	group.GroupId = 1
	return group, nil
}
func (r *DummyRepository) GetAllGroups() ([]models.Group, error) {
	groups := make([]models.Group, sutTableRuns)
	for i := 0; i < sutTableRuns; i++ {
		g := validGroup
		g.GroupId = i + 1
		groups[i] = g
	}
	return groups, nil
}
func (r *DummyRepository) GetGroupById(groupId int) (models.Group, error) {
	ctime := time.Now().UTC().UnixMilli()

	var group models.Group
	group.GroupId = groupId
	group.CreatedAt = ctime
	group.CreatorId = 1
	group.Description = "It's a Dummy Group"
	group.Title = "Dummy Group"
	group.UpdatedAt = ctime

	return group, nil
}

func (r *DummyRepository) UpdateGroup(group models.Group) (models.Group, error) {
	return group, nil
}
func (r *DummyRepository) DeleteGroup(groupId int) error {
	return nil
}
func (r *DummyRepository) DeleteGroupUser(groupUser int) error {
	return nil
}
func (r *DummyRepository) DeleteAllGroups() error {
	return nil
}

// GroupUser
func (r *DummyRepository) CreateGroupUser(groupUser models.GroupUser) (models.GroupUser, error) {
	return validGroupUser, nil
}
func (r *DummyRepository) GetGroupUser(groupUserId int) (models.GroupUser, error) {
	groupUser := validGroupUser
	groupUser.GroupUserId = groupUserId
	return groupUser, nil
}
func (r *DummyRepository) GetGroupUsersByUserId(userId int) ([]models.GroupUser, error) {
	var groupUsers []models.GroupUser
	for i := 0; i < rand.Intn(6); i++ {
		groupUser := validGroupUser
		groupUser.UserId = userId
		groupUsers = append(groupUsers, groupUser)
	}

	return groupUsers, nil
}
func (r *DummyRepository) GetGroupUsersByGroupId(groupId int) ([]models.GroupUser, error) {
	var groupUsers []models.GroupUser

	for i := 0; i < rand.Intn(6); i++ {
		groupUser := validGroupUser
		groupUser.GroupId = groupId
		groupUsers = append(groupUsers, groupUser)
	}

	return groupUsers, nil
}
func (r *DummyRepository) DeleteGroupUsersByUserId(userId int) error {
	return errors.New("dummy deleted groupusers by userid")
}
func (r *DummyRepository) DeleteGroupUserByGroupId(groupId int) error {
	return nil
}
func (r *DummyRepository) DeleteGroupUserByGroupIdAndUserId(groupId, userId int) error {
	return nil
}
func (r *DummyRepository) DeleteAllGroupUsers() error {
	return nil
}

// Notifications
func (r *DummyRepository) CreateNotification(notification models.Notification) (models.Notification, error) {
	return notification, nil
}
func (r *DummyRepository) GetNotificationById(notificationId int) (models.Notification, error) {
	ctime := time.Now().UTC().UnixMilli()

	notification := models.Notification{
		NotificationId:   notificationId,
		CreatedAt:        ctime,
		NotificationType: "dummy type",
		ObjectId:         1,
		SenderId:         1,
		Status:           "dummy status",
		TargetId:         2,
		UpdatedAt:        ctime,
	}

	return notification, nil
}
func (r *DummyRepository) GetNotificationsByTargetId(targetId int) ([]models.Notification, error) {

	var notifications []models.Notification
	for i := 0; i < sutTableRuns; i++ {
		n := validNotification
		n.NotificationId = i + 1
		n.TargetId = targetId
		notifications = append(notifications, n)
	}
	return notifications, nil
}
func (r *DummyRepository) UpdateNotification(notification models.Notification) (models.Notification, error) {
	ctime := time.Now().UTC().UnixMilli()

	notification.NotificationType = "dummy type update"
	notification.Status = "dummy status updated"
	notification.UpdatedAt = ctime

	return notification, nil
}
func (r *DummyRepository) DeleteNotificationById(notificationId int) error {
	return nil
}

// Message from Matt
// Please do not touch this file

package repo

import (
	"log"
	"socialnetwork/models"
	"socialnetwork/sqlite"
	comments "socialnetwork/sqlite/COMMENTS"
	events "socialnetwork/sqlite/EVENTS"
	event_users "socialnetwork/sqlite/EVENT_USERS"
	groups "socialnetwork/sqlite/GROUPS"
	group_users "socialnetwork/sqlite/GROUP_USERS"
	messages "socialnetwork/sqlite/MESSAGES"
	notifications "socialnetwork/sqlite/NOTIFICATIONS"
	posts "socialnetwork/sqlite/POSTS"
	post_users "socialnetwork/sqlite/POST_USERS"
	users "socialnetwork/sqlite/USERS"
	user_users "socialnetwork/sqlite/USER_USERS"
	"socialnetwork/transport"

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

// Home
func (r *SQLiteRepository) GetHomeDataForUser(userId int) (transport.HomeModel, error) {
	return sqlite.GetHomeDataForUser(r.identityDb, r.businessDb, userId)
}

// Profile
func (r *SQLiteRepository) GetProfileDataForUser(userId int) (transport.ProfileModel, error) {
	return sqlite.GetProfileDataForUser(r.identityDb, r.businessDb, userId)
}
func (r *SQLiteRepository) UpdateIsPublic(userId int, isPublic bool) error {
	return sqlite.UpdateIsPublic(r.identityDb, userId, isPublic)
}

// Users
func (r *SQLiteRepository) CreateUser(user models.User) (models.User, error) {
	return users.CreateUser(r.identityDb, user)
}
func (r *SQLiteRepository) GetAllUsers() ([]models.User, error) {
	return users.GetAllUsers(r.identityDb)
}
func (r *SQLiteRepository) GetAllUsersTransport() ([]transport.UserTransport, error) {
	return sqlite.GetAllUsersTransport(r.identityDb)
}
func (r *SQLiteRepository) GetUsersByPublic() ([]models.User, error) {
	return users.GetUsersByPublic(r.identityDb)
}
func (r *SQLiteRepository) GetUserById(userId int) (models.User, error) {
	return users.GetUserById(r.identityDb, userId)
}

// func (r *SQLiteRepository) GetUserByEmail(email string) (models.User, error) {
// 	return sqlite.GetUserByEmail(r.identityDb, email)
// }
// func (r *SQLiteRepository) GetUserByUsername(username string) (models.User, error) {
// 	return sqlite.GetUserByUsername(r.identityDb, username)
// }

func (r *SQLiteRepository) GetUserByUsernameOrEmail(usernameOrEmail string) (models.User, error) {
	return users.GetUserByUsernameOrEmail(r.identityDb, usernameOrEmail)
}
func (r *SQLiteRepository) UpdateUser(user models.User) (models.User, error) {
	return users.UpdateUser(r.identityDb, user)
}
func (r *SQLiteRepository) DeleteUserById(userId int) error {
	return users.DeleteUserById(r.identityDb, userId)
}
func (r *SQLiteRepository) DeleteAllUsers() error {
	return users.DeleteAllUsers(r.identityDb)
}

// UserUser
func (r *SQLiteRepository) CreateUserUser(userUser models.UserUser) (models.UserUser, error) {
	return user_users.CreateUserUser(r.businessDb, userUser)
}
func (r *SQLiteRepository) GetUserUsersBySubjectId(subjectId int) ([]models.UserUser, error) {
	return user_users.GetUserUsersBySubjectId(r.businessDb, subjectId)
}
func (r *SQLiteRepository) GetUserUsersByFollowerId(followerId int) ([]models.UserUser, error) {
	return user_users.GetUserUsersByFollowerId(r.businessDb, followerId)
}
func (r *SQLiteRepository) GetUserUserByFollowerIdAndSubjectId(followerId, subjectId int) (models.UserUser, error) {
	return user_users.GetUserUserByFollowerIdAndSubjectId(r.businessDb, followerId, subjectId)
}

func (r *SQLiteRepository) DeleteUserUsersByFollowerId(followerId int) error {
	return user_users.DeleteUserUsersByFollowerId(r.businessDb, followerId)
}
func (r *SQLiteRepository) DeleteUserUsersBySubjectId(subjectId int) error {
	return user_users.DeleteUserUsersByFollowerId(r.businessDb, subjectId)
}
func (r *SQLiteRepository) DeleteUserUserBySubjectIdAndFollowerId(subjectId, followerId int) error {
	return user_users.DeleteUserUserBySubjectIdAndFollowerId(r.businessDb, subjectId, followerId)
}

// Post
func (r *SQLiteRepository) CreatePost(post models.Post) (models.Post, error) {
	return posts.CreatePost(r.businessDb, post)
}
func (r *SQLiteRepository) GetAllPosts() ([]models.Post, error) {
	return posts.GetAllPosts(r.businessDb)
}
func (r *SQLiteRepository) GetPostById(postId int) (models.Post, error) {
	return posts.GetPostById(r.businessDb, postId)
}
func (r *SQLiteRepository) GetPostsByGroupId(groupId int) ([]models.Post, error) {
	return posts.GetPostsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	return posts.GetPostsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetPostsByPrivacy(privacy string) ([]models.Post, error) {
	return posts.GetPostsByPrivacy(r.businessDb, privacy)
}
func (r *SQLiteRepository) DeletePostById(postId int) error {
	return posts.DeletePostById(r.businessDb, postId)
}
func (r *SQLiteRepository) UpdatePost(post models.Post) (models.Post, error) {
	return posts.UpdatePost(r.businessDb, post)
}
func (r *SQLiteRepository) DeletePostByGroupId(groupId int) error {
	return posts.DeletePostByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeletePostsByUserId(userId int) error {
	return posts.DeletePostsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteAllPosts() error {
	return posts.DeleteAllPosts(r.businessDb)
}

// Post_Users
func (r *SQLiteRepository) CreatePostUser(postUser models.PostUser) (models.PostUser, error) {
	return post_users.CreatePostUser(r.businessDb, postUser)
}
func (r *SQLiteRepository) GetPostUsersByUserId(userId int) ([]models.PostUser, error) {
	return post_users.GetPostUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetPostUsersByPostId(postId int) ([]models.PostUser, error) {
	return post_users.GetPostUsersByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) DeletePostUsersByUserId(userId int) error {
	return post_users.DeletePostUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeletePostUsersByPostId(postId int) error {
	return post_users.DeletePostUsersByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) DeletePostUserByPostIdAndUserId(postId, userId int) error {
	return post_users.DeletePostUserByPostIdAndUserId(r.businessDb, postId, userId)
}
func (r *SQLiteRepository) DeleteAllPostUsers() error {
	return post_users.DeleteAllPostUsers(r.businessDb)
}

// Comments
func (r *SQLiteRepository) CreateComment(comment models.Comment) (models.Comment, error) {
	return comments.CreateComment(r.businessDb, comment)
}
func (r *SQLiteRepository) GetAllComments() ([]models.Comment, error) {
	return comments.GetAllComments(r.businessDb)
}
func (r *SQLiteRepository) GetCommentById(commentId int) (models.Comment, error) {
	return comments.GetCommentById(r.businessDb, commentId)
}

//	func (r *SQLiteRepository) GetCommentsByGroupId(groupId int) ([]models.Comment, error) {
//		return comments.GetCommentsByGroupId(r.businessDb, groupId)
//	}
func (r *SQLiteRepository) GetCommentsByUserId(userId int) ([]models.Comment, error) {
	return comments.GetCommentsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetCommentsByPostId(postId int) ([]models.Comment, error) {
	return comments.GetCommentsByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	return comments.UpdateComment(r.businessDb, comment)
}
func (r *SQLiteRepository) DeleteCommentById(commentId int) error {
	return comments.DeleteCommentById(r.businessDb, commentId)
}
func (r *SQLiteRepository) DeleteCommentsByGroupId(groupId int) error {
	return comments.DeleteCommentsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteCommentsByUserId(userId int) error {
	return comments.DeleteCommentsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteCommentsByPostId(postId int) error {
	return comments.DeleteCommentsByPostId(r.businessDb, postId)
}
func (r *SQLiteRepository) DeleteAllComments() error {
	return comments.DeleteAllComments(r.businessDb)
}

// Event
func (r *SQLiteRepository) CreateEvent(event models.Event) (models.Event, error) {
	return events.CreateEvent(r.businessDb, event)
}
func (r *SQLiteRepository) GetAllEvents() ([]models.Event, error) {
	return events.GetAllEvents(r.businessDb)
}
func (r *SQLiteRepository) GetEventById(eventId int) (models.Event, error) {
	return events.GetEventById(r.businessDb, eventId)
}
func (r *SQLiteRepository) GetEventsByGroupId(groupId int) ([]models.Event, error) {
	return events.GetEventsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) GetEventsByUserId(userId int) ([]models.Event, error) {
	return events.GetEventsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) UpdateEvent(event models.Event) (models.Event, error) {
	return events.UpdateEvent(r.businessDb, event)
}
func (r *SQLiteRepository) DeleteEventById(eventId int) error {
	return events.DeleteEventById(r.businessDb, eventId)
}
func (r *SQLiteRepository) DeleteEventsByGroupId(groupId int) error {
	return events.DeleteEventsByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteEventsByUserId(userId int) error {
	return events.DeleteEventsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteAllEvents() error {
	return events.DeleteAllEvents(r.businessDb)
}

// EventUser
func (r *SQLiteRepository) CreateEventUser(eventUser models.EventUser) (models.EventUser, error) {
	return event_users.CreateEventUser(r.businessDb, eventUser)
}
func (r *SQLiteRepository) GetEventUsersByUserId(userId int) ([]models.EventUser, error) {
	return event_users.GetEventUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetEventUsersByEventId(eventId int) ([]models.EventUser, error) {
	return event_users.GetEventUsersByUserId(r.businessDb, eventId)
}
func (r *SQLiteRepository) DeleteEventUsersByUserId(userId int) error {
	return event_users.DeleteEventUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteEventUsersByEventId(eventId int) error {
	return event_users.DeleteEventUsersByEventId(r.businessDb, eventId)
}
func (r *SQLiteRepository) DeleteEventUserByEventIdAndUserId(eventId, userId int) error {
	return event_users.DeleteEventUserByEventIdAndUserId(r.businessDb, userId, eventId)
}
func (r *SQLiteRepository) DeleteAllEventUsers() error {
	return event_users.DeleteAllEventUsers(r.businessDb)
}

// Message
func (r *SQLiteRepository) CreateMessage(message models.Message) (models.Message, error) {
	return messages.CreateMessage(r.businessDb, message)
}
func (r *SQLiteRepository) GetAllMessages() ([]models.Message, error) {
	return messages.GetAllMessages(r.businessDb)
}
func (r *SQLiteRepository) GetMessagesByMessageTypeandTargetId(messageType string, targetId int) ([]models.Message, error) {
	return messages.GetMessagesByMessageTypeandTargetId(r.businessDb, messageType, targetId)
}

func (r *SQLiteRepository) GetMessageById(messageId int) (models.Message, error) {
	return messages.GetMessageById(r.businessDb, messageId)
}
func (r *SQLiteRepository) GetMessagesBySenderAndTargetIDs(senderId, targetId int) ([]models.Message, error) {
	return messages.GetMessagesBySenderAndTargetIds(r.businessDb, senderId, targetId)
}
func (r *SQLiteRepository) UpdateMessage(message models.Message) (models.Message, error) {
	return messages.UpdateMessage(r.businessDb, message)
}
func (r *SQLiteRepository) DeleteMessagesByType(messageType string) error {
	return messages.DeleteMessagesByType(r.businessDb, messageType)
}
func (r *SQLiteRepository) DeleteMessageById(messageId int) error {
	return messages.DeleteMessageById(r.businessDb, messageId)
}
func (r *SQLiteRepository) DeleteMessagesBySenderId(senderId int) error {
	return messages.DeleteMessagesBySenderId(r.businessDb, senderId)
}
func (r *SQLiteRepository) DeleteMessagesByTargetId(targetId int) error {
	return messages.DeleteMessagesByTargetId(r.businessDb, targetId)
}
func (r *SQLiteRepository) DeleteAllMessages() error {
	return messages.DeleteAllMessages(r.businessDb)
}

// Group
func (r *SQLiteRepository) CreateGroup(group models.Group) (models.Group, error) {
	return groups.CreateGroup(r.businessDb, group)
}

func (r *SQLiteRepository) GetAllGroups() ([]models.Group, error) {
	return groups.GetAllGroups(r.businessDb)
}
func (r *SQLiteRepository) GetGroupById(groupId int) (models.Group, error) {
	return groups.GetGroupById(r.businessDb, groupId)
}
func (r *SQLiteRepository) UpdateGroup(group models.Group) (models.Group, error) {
	return groups.UpdateGroup(r.businessDb, group)
}
func (r *SQLiteRepository) DeleteGroup(groupId int) error {
	return groups.DeleteGroup(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteAllGroups() error {
	return groups.DeleteAllGroups(r.businessDb)
}

// Group_User
func (r *SQLiteRepository) CreateGroupUser(groupUser models.GroupUser) (models.GroupUser, error) {
	return group_users.CreateGroupUser(r.businessDb, groupUser)
}
func (r *SQLiteRepository) GetGroupUser(groupUserId int) (models.GroupUser, error) {
	return group_users.GetGroupUser(r.businessDb, groupUserId)
}
func (r *SQLiteRepository) GetGroupUsersByUserId(userId int) ([]models.GroupUser, error) {
	return group_users.GetGroupUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) GetGroupUsersByGroupId(groupId int) ([]models.GroupUser, error) {
	return group_users.GetGroupUsersByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteGroupUsersByUserId(userId int) error {
	return group_users.DeleteGroupUsersByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) DeleteGroupUserByGroupId(groupId int) error {
	return group_users.DeleteGroupUserByGroupId(r.businessDb, groupId)
}
func (r *SQLiteRepository) DeleteGroupUserByGroupIdAndUserId(groupId, userId int) error {
	return group_users.DeleteGroupUserByGroupIdAndUserId(r.businessDb, groupId, userId)
}
func (r *SQLiteRepository) DeleteGroupUser(groupUserId int) error {
	return group_users.DeleteGroupUser(r.businessDb, groupUserId)
}

func (r *SQLiteRepository) DeleteAllGroupUsers() error {
	return group_users.DeleteAllGroupUsers(r.businessDb)
}

// Notification
func (r *SQLiteRepository) CreateNotification(notification models.Notification) (models.Notification, error) {
	return notifications.CreateNotification(r.businessDb, notification)
}
func (r *SQLiteRepository) GetNotificationById(notificationId int) (models.Notification, error) {
	return notifications.GetNotificationById(r.businessDb, notificationId)
}

func (r *SQLiteRepository) GetNotificationsByUserId(userId int) ([]models.Notification, error) {
	return notifications.GetNotificationsByUserId(r.businessDb, userId)
}
func (r *SQLiteRepository) UpdateNotification(notification models.Notification) (models.Notification, error) {
	return notifications.UpdateNotification(r.businessDb, notification)
}
func (r *SQLiteRepository) DeleteNotificationById(notificationId int) error {
	return notifications.DeleteNotificationById(r.businessDb, notificationId)
}

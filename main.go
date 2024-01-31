package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"socialnetwork/api"
	"socialnetwork/auth"
	"socialnetwork/repo"
	"socialnetwork/router"
	"socialnetwork/ui"
)

const (
	addr = ":8080"
)

func main() {
	// Setup serve mux
	mux := http.NewServeMux()
	// Host static files
	serveStaticFiles(mux)
	// Add handlers to router
	setupRouter(mux)
	// Listen and serve
	fmt.Printf("server listening at address %s...\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func setupRouter(mux *http.ServeMux) {
	rt := router.NewRouter()
	addApiHandlers(rt)
	addUIHandlers(rt)
	mux.Handle("/", rt)
}

func addApiHandlers(rt *router.Router) {
	r := repo.NewDummyRepository()
	loginHandler := auth.NewLoginHandler(r)
	logoutHandler := auth.NewLogoutHandler(r)
	registrationHandler := auth.NewRegistrationHandler(r)
	usersHandler := api.NewUsersHandler(r)
	userByIdHandler := api.NewUserByIdHandler(r)
	postsHandler := api.NewPostsHandler(r)
	postByIdHandler := api.NewPostByIdHandler(r)
	userPostsHandler := api.NewUserPostsHandler(r)
	postByPrivacyHandler := api.NewPostByPrivacyHandler(r)
	commentsHandler := api.NewCommentsHandler(r)
	commentByIdHandler := api.NewCommentByIdHandler(r)
	commentByPostIdHandler := api.NewCommentsByPostIdHandler(r)
	eventsHandler := api.NewEventsHandler(r)
	eventByIdHandler := api.NewEventByIdHandler(r)
	eventsByGroupIdHandler := api.NewEventsByGroupIdHandler(r)
	eventUsersHandler := api.NewEventUsersHandler(r)
	eventUsersByEventIdHandler := api.NewEventUsersByEventIdHandler(r)
	eventUsersByEventIdAndUserIdHandler := api.NewEventUserByEventIdAndUserIdHandler(r)
	messagesHandler := api.NewMessagesHandler(r)
	messageByIdHandler := api.NewMessageByIdHandler(r)
	messagesByGroupIdHandler := api.NewMessagesByGroupIdHandler(r)
	groupsHandler := api.NewGroupsHandler(r)
	groupByIdHandler := api.NewGroupByIdHandler(r)
	groupUsersHandler := api.NewGroupUsersHandler(r)
	groupUserByIdHandler := api.NewGroupUserByIdHandler(r)
	groupUserByGroupIdAndUserIdHandler := api.NewGroupUserByGroupIdAndUserIdHandler(r)
	groupPostsHandler := api.NewGroupPostsHandler(r)
	notificationsHandler := api.NewNotificationsHandler(r)
	notificationByIdHandler := api.NewNotificationByIdHandler(r)

	// Auth Handlers
	rt.AddHandler(regexp.MustCompile(`^/auth/login$`), loginHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/logoutn$`), logoutHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/registration$`), registrationHandler)

	// User Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/users$`), usersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/{userId}$`), userByIdHandler)
	// Post Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/posts$`), postsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/{postId}$`), postByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/user/{userId}/posts$`), userPostsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/privacy/{privacyStatus}$`), postByPrivacyHandler)
	// Comment Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/comments$`), commentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/comments/{commentId}$`), commentByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/{postId}/comments$`), commentByPostIdHandler)
	// Event Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/events$`), eventsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}$`), eventByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/events$`), eventsByGroupIdHandler)
	// EventUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/eventUsers$`), eventUsersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}/eventUsers$`), eventUsersByEventIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}/eventUsers/users/{userId}$`), eventUsersByEventIdAndUserIdHandler)
	// Message Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/messages$`), messagesHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/messages/{messageId$`), messageByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/messages$`), messagesByGroupIdHandler)
	// Group Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groups$`), groupsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId$`), groupByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/posts$`), groupPostsHandler)
	// GroupUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groupUsers$`), groupUsersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groupUsers/{groupUserId$`), groupUserByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/eventUsers/users/{userId}$`), groupUserByGroupIdAndUserIdHandler)
	// Notification Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/notifications$`), notificationsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/notifications/{notificationId}$`), notificationByIdHandler)
}

func addUIHandlers(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/$`), ui.NewDummyPageHandler())
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

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
	"socialnetwork/ws"
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
	addWSHandler(rt)
	addUIHandlers(rt)
	mux.Handle("/", rt)
}

func addApiHandlers(rt *router.Router) {
	// r := repo.NewDummyRepository()
	r := repo.NewSQLiteRepository()
	loginHandler := auth.NewLoginHandler(r)
	logoutHandler := auth.NewLogoutHandler(r)
	registrationHandler := auth.NewRegistrationHandler(r)
	usersHandler := api.NewUsersHandler(r)
	userByIdHandler := api.NewUserByIdHandler(r)

	postsHandler := api.NewPostsHandler(r)
	postByIdHandler := api.NewPostByIdHandler(r)
	// postUsersByPostIdHandler := api.NewPostUsersByPostIdHandler(r)
	// userPostsHandler := api.NewUserPostsHandler(r)
	// postByPrivacyHandler := api.NewPostsByPrivacyHandler(r)
	// postUserHandler := api.NewPostUsersHandler(r)
	commentsHandler := api.NewCommentsHandler(r)
	commentByIdHandler := api.NewCommentByIdHandler(r)
	commentByPostIdHandler := api.NewCommentsByPostIdHandler(r)
	eventsHandler := api.NewEventsHandler(r)
	eventByIdHandler := api.NewEventByIdHandler(r)
	eventsByGroupIdHandler := api.NewEventsByGroupIdHandler(r)
	eventsByUserIdHandler := api.NewEventsByUserIdHandler(r)
	eventUsersHandler := api.NewEventUsersHandler(r)
	// eventUsersByEventIdHandler := api.NewEventUsersByEventIdHandler(r)
	// eventUsersByEventIdAndUserIdHandler := api.NewEventUserByEventIdAndUserIdHandler(r)
	// messagesHandler := api.NewMessagesHandler(r)
	// messageByIdHandler := api.NewMessageByIdHandler(r)
	groupsHandler := api.NewGroupsHandler(r)
	groupByIdHandler := api.NewGroupByIdHandler(r)
	groupUsersHandler := api.NewGroupUsersHandler(r)
	groupUsersByGroupIdHandler := api.NewGroupUsersByGroupIdHandler(r)
	groupUsersByUserIdHandler := api.NewGroupUsersByUserIdHandler(r)
	// groupUserByIdHandler := api.NewGroupUserByIdHandler(r)
	// groupUserByGroupIdAndUserIdHandler := api.NewGroupUserByGroupIdAndUserIdHandler(r)
	postsByGroupIdHandler := api.NewPostsByGroupIdHandler(r)
	messagesByTypeAndTargetIdHandler := api.NewMessagesByTypeAndTargetIdHandler(r)
	notificationsHandler := api.NewNotificationsHandler(r)
	notificationByIdHandler := api.NewNotificationByIdHandler(r)
	notificationsByUserIdHandler := api.NewNotificationsByUserIdHandler(r)
	userUsersHandler := api.NewUserUsersHandler(r)
	userUsersByFollowerIdHandler := api.NewUserUsersByFollowerIdHandler(r)
	userUsersBySubjectIdHandler := api.NewUserUsersBySubjectIdHandler(r)
	UserUserBySubjectIdAndFollowerIdHandler := api.NewUserUserBySubjectIdAndFollowerIdHandler(r)

	userIdHandler := api.NewUserIdHandler(r)
	homeHandler := api.NewHomeHandler(r)
	profileHandler := api.NewProfileHandler(r)
	privacyHandler := api.NewPrivacyHandler(r)

	// Auth Handlers
	rt.AddHandler(regexp.MustCompile(`^/auth/login$`), loginHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/logout$`), logoutHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/registration$`), registrationHandler)

	// // User Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/users$`), usersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+$`), userByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/groupUsers$`), groupUsersByUserIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/notifications$`), notificationsByUserIdHandler)

	// rt.AddHandler(regexp.MustCompile(`^/api/users/privacy/public$`), usersByPublicHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/posts/[0-9]+$`), postUserByPostIdAndUserIdHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+$`), postUsersByPostIdHandler)
	// // Post Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/posts$`), postsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/[0-9]+$`), postByIdHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/user/{userId}/posts$`), userPostsHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/posts/privacy/{privacyStatus}$`), postByPrivacyHandler)
	// // PostUser Handlers
	// rt.AddHandler(regexp.MustCompile(`^/api/postUsers/$`), postUserHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/postUsers/users/{userId$`), postUserByUserIdHandler)
	// // Comment Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/comments$`), commentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/comments/[0-9]+$`), commentByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/[0-9]+/comments$`), commentByPostIdHandler)
	// // Event Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/events$`), eventsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/[0-9]+$`), eventByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/events$`), eventsByUserIdHandler)

	// // EventUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/eventUsers$`), eventUsersHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}/eventUsers$`), eventUsersByEventIdHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}/eventUsers/users/{userId}$`), eventUsersByEventIdAndUserIdHandler)
	// // Message Handlers
	// rt.AddHandler(regexp.MustCompile(`^/api/messages$`), messagesHandler)

	// // Group Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groups$`), groupsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+$`), groupByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/posts$`), postsByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/groupUsers$`), groupUsersByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/messages$`), messagesByTypeAndTargetIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/events$`), eventsByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/events/[0-9]+$`), eventByIdHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/{groupUserId$`), groupUserByIdHandler)

	// // GroupUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groupUsers$`), groupUsersHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/{groupUserId$`), groupUserByIdHandler)
	// rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/users/{userId}$`), groupUserByGroupIdAndUserIdHandler)

	// // Notification Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/notifications$`), notificationsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/notifications/[0-9]+$`), notificationByIdHandler)
	// UserUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/userUsers$`), userUsersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/userUsers$`), userUsersByFollowerIdHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/followerUserUsers$`), userUsersBySubjectIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/userUsers/[0-9]+$`), UserUserBySubjectIdAndFollowerIdHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/userId$`), userIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/home$`), homeHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/profile/[0-9]+$`), profileHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/profile/privacy$`), privacyHandler)
}

func addWSHandler(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/ws$`), ws.NewWebSocketHandler(repo.NewDummyRepository()))
}

func addUIHandlers(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/$`), ui.NewPageHandler())
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

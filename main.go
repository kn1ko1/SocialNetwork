package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"socialnetwork/api"
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
	usersHandler := api.NewUsersHandler(r)
	userByIdHandler := api.NewUserByIdHandler(r)
	postsHandler := api.NewPostsHandler(r)
	postByIdHandler := api.NewPostByIdHandler(r)
	userPostsHandler := api.NewUserPostsHandler(r)
	commentsHandler := api.NewCommentsHandler(r)
	commentByIdHandler := api.NewCommentByIdHandler(r)
	commentByPostIdHandler := api.NewCommentsByPostIdHandler(r)
	eventsHandler := api.NewEventsHandler(r)
	eventByIdHandler := api.NewEventByIdHandler(r)
	eventsByGroupIdHandler := api.NewEventsByGroupIdHandler(r)
	messagesHandler := api.NewMessagesHandler(r)
	messageByIdHandler := api.NewMessageByIdHandler(r)
	messagesByGroupIdHandler := api.NewMessagesByGroupIdHandler(r)
	groupsHandler := api.NewGroupsHandler(r)
	groupUsersHandler := api.NewGroupUsersHandler(r)
	groupPostsHandler := api.NewGroupPostsHandler(r)
	notificationsHandler := api.NewNotificationsHandler(r)
	notificationByIdHandler := api.NewNotificationByIdHandler(r)

	// User Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/users$`), usersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/{userId}$`), userByIdHandler)
	// Post Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/posts$`), postsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/{postId}$`), postByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/user/{userId}/posts$`), userPostsHandler)
	// Comment Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/comments$`), commentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/comments/{commentId}$`), commentByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/{postId}/comments$`), commentByPostIdHandler)
	// Event Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/events$`), eventsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/{eventId}$`), eventByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/events$`), eventsByGroupIdHandler)
	// EventGroup Handlers
	// Message Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/messages$`), messagesHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/messages$`), messageByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/messages$`), messagesByGroupIdHandler)
	// Group Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groups$`), groupsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/{groupId}/posts$`), groupPostsHandler)
	// GroupUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groupUsers$`), groupUsersHandler)
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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	dbUtils "socialnetwork/Database/databaseUtils"
	"socialnetwork/Server/api"
	"socialnetwork/Server/auth"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/router"
	"socialnetwork/Server/ui"
	"socialnetwork/Server/ws"
)

const (
	addr = ":8080"
)

var (
	socketGroupManager *ws.SocketGroupManager
)

func initDatabases() {

	dbUtils.InitIdentityDatabase()
	dbUtils.InitBusinessDatabase()

	ui.InitTemplates()

	socketGroupManager = ws.NewSocketGroupManager()
	socketGroupManager.Start()
}

func main() {
	initDatabases()

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
	addImageHandlers(rt)
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
	postsByUserIdHandler := api.NewPostsByUserIdHandler(r)
	usersByFollowerIdHandler := api.NewUsersByFollowerIdHandler(r)
	usersBySubjectIdHandler := api.NewUsersBySubjectIdHandler(r)

	postsHandler := api.NewPostsHandler(r)
	postByIdHandler := api.NewPostByIdHandler(r)
	postsByGroupIdHandler := api.NewPostsByGroupIdHandler(r)
	postsByGroupIdWithCommentsHandler := api.NewPostsByGroupIdWithCommentsHandler(r)
	PostsPublicWithCommentsHandler := api.NewPostsPublicWithCommentsHandler(r)
	postsPrivateWithCommentsHandler := api.NewPostsPrivateWithCommentsHandler(r)
	postsAlmostPrivateWithCommentsHandler := api.NewPostsAlmostPrivateWithCommentsHandler(r)

	commentsHandler := api.NewCommentsHandler(r)
	commentByIdHandler := api.NewCommentByIdHandler(r)
	commentByPostIdHandler := api.NewCommentsByPostIdHandler(r)

	eventsHandler := api.NewEventsHandler(r)
	eventByIdHandler := api.NewEventByIdHandler(r)
	eventsByGroupIdHandler := api.NewEventsByGroupIdHandler(r)
	eventsByUserIdHandler := api.NewEventsByUserIdHandler(r)
	eventUsersHandler := api.NewEventUsersHandler(r)

	groupsHandler := api.NewGroupsHandler(r)
	groupByIdHandler := api.NewGroupByIdHandler(r)
	groupsByUserIdHandler := api.NewGroupsByUserIdHandler(r)
	groupUsersHandler := api.NewGroupUsersHandler(r)
	groupUsersByGroupIdHandler := api.NewGroupUsersByGroupIdHandler(r)
	groupUsersByUserIdHandler := api.NewGroupUsersByUserIdHandler(r)

	messagesBySenderAndTargetIDHandler := api.NewMessagesBySenderAndTargetIDHandler(r)
	messagesByTypeAndTargetIdHandler := api.NewMessagesByTypeAndTargetIdHandler(r)

	notificationsHandler := api.NewNotificationsHandler(r)
	notificationByIdHandler := api.NewNotificationByIdHandler(r)
	notificationsByUserIdHandler := api.NewNotificationsByUserIdHandler(r)

	userUsersHandler := api.NewUserUsersHandler(r)
	userUsersByFollowerIdHandler := api.NewUserUsersByFollowerIdHandler(r)
	userUsersBySubjectIdHandler := api.NewUserUsersBySubjectIdHandler(r)
	UserUserBySubjectIdAndFollowerIdHandler := api.NewUserUserBySubjectIdAndFollowerIdHandler(r)

	userIdHandler := api.NewUserIdHandler(r)
	privacyHandler := api.NewPrivacyHandler(r)

	// Auth Handlers
	rt.AddHandler(regexp.MustCompile(`^/auth/login$`), loginHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/logout$`), logoutHandler)
	rt.AddHandler(regexp.MustCompile(`^/auth/registration$`), registrationHandler)

	// // User Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/users$`), usersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+$`), userByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/posts$`), postsByUserIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/groupUsers$`), groupUsersByUserIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/notifications$`), notificationsByUserIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/followerUsers$`), usersByFollowerIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/followedUsers$`), usersBySubjectIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/messages/[0-9]+$`), messagesBySenderAndTargetIDHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/posts$`), postsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/[0-9]+$`), postByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/posts$`), postsByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/posts/withComments$`), postsByGroupIdWithCommentsHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/posts/public/withComments$`), PostsPublicWithCommentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/private/withComments$`), postsPrivateWithCommentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/almostPrivate/withComments$`), postsAlmostPrivateWithCommentsHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/comments$`), commentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/comments/[0-9]+$`), commentByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts/[0-9]+/comments$`), commentByPostIdHandler)
	// // Event Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/events$`), eventsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events/[0-9]+$`), eventByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/events$`), eventsByUserIdHandler)

	// // EventUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/eventUsers$`), eventUsersHandler)

	// // Group Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groups$`), groupsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+$`), groupByIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/groups$`), groupsByUserIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/groupUsers$`), groupUsersByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/messages$`), messagesByTypeAndTargetIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/events$`), eventsByGroupIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups/[0-9]+/events/[0-9]+$`), eventByIdHandler)

	// // GroupUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/groupUsers$`), groupUsersHandler)

	// // Notification Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/notifications$`), notificationsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/notifications/[0-9]+$`), notificationByIdHandler)
	// UserUser Handlers
	rt.AddHandler(regexp.MustCompile(`^/api/userUsers$`), userUsersHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/userUsers$`), userUsersByFollowerIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/followerUserUsers$`), userUsersBySubjectIdHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/users/[0-9]+/userUsers/[0-9]+$`), UserUserBySubjectIdAndFollowerIdHandler)

	rt.AddHandler(regexp.MustCompile(`^/api/userId$`), userIdHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/profile/privacy$`), privacyHandler)
}

func addWSHandler(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/ws$`), ws.NewWebSocketHandler(repo.NewSQLiteRepository()))
}

func addUIHandlers(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/$`), ui.NewPageHandler())
}

func addImageHandlers(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/uploads/.*$`), ui.NewImageHandler())
}

func serveStaticFiles(mux *http.ServeMux) {
	workDir, _ := os.Getwd()
	staticDir := filepath.Join(workDir, "..", "App", "static") // Go up one level, then into 'App/static'
	fsRoot := http.Dir(staticDir)
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", enableCors(http.StripPrefix("/static/", fs)))

	log.Println("[Server/main.go] Serving static files from:", staticDir) // Log the path for debugging
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // or specify a specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

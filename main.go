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
	groupsHandler := api.NewGroupsHandler(r)
	postsHandler := api.NewPostsHandler(r)
	commentsHandler := api.NewCommentsHandler(r)
	eventsHandler := api.NewEventsHandler(r)
	notificationsHandler := api.NewNotificationsHandler(r)
	messagesHandler := api.NewMessagesHandler(r)
	rt.AddHandler(regexp.MustCompile(`^/api/users$`), usersHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/groups$`), groupsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/posts$`), postsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/comments$`), commentsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/events$`), eventsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/notifications$`), notificationsHandler)
	rt.AddHandler(regexp.MustCompile(`^/api/messages$`), messagesHandler)
}

func addUIHandlers(rt *router.Router) {
	rt.AddHandler(regexp.MustCompile(`^/$`), ui.NewDummyPageHandler())
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

package ui

import (
	"fmt"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/posts
// Allowed methods: GET, POST

type PostsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsHandler(r repo.IRepository) *PostsHandler {
	return &PostsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get Session Cookie
	c, err := r.Cookie("Session")
	if err != nil {
		// Log Error
		log.Println(err.Error())
		// Return HTTP Status Unauthorized
		//
		// N.B. for simplicity of the example, we are simply returning
		// an HTTP error. In the actual project, probably a JSON payload.
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Authenticate Session Cookie - user variable discarded because user struct not used here...
	_, err = auth.AuthenticateSessionCookie(c)
	if err != nil {
		// Same error as above - maker of request is unauthorized
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// HTTP GET logic
	case http.MethodGet:
		// Not Implemented
		h.get(w, r)
		return
	// HTTP POST logic
	case http.MethodPost:
		h.post(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Separate out HTTP methods for clean separation of concerns
// N.B. Use lowercase names, i.e. "post", "get", etc. for correct encapsulation
func (h *PostsHandler) post(w http.ResponseWriter, r *http.Request) {
	// Read the JSON body of the request OR parse form data to get the post
	// (we are gonna do both, using headers correctly we can separate UI logic from API logic)
	//
	// Again, ommitted here for sake of example. We just assume this is what user is trying to post:
	post := models.Post{UserId: 1, Title: "Example"}
	// Self-contained Validation pipeline method
	// If this fails - Bad Request
	// err := post.Validate()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// 	return
	// }
	// Handler uses its Repo instance to add the post to the DB
	// The Repo instance itself is responsible for its own
	// data access layer (DAL) implementation. e.g. SQLite, MySQL, etc.
	//
	// If this fails - internal server error
	result, err := h.Repo.CreatePost(post)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// Instead of just printing here, we would then put the result in a JSON
	// payload, to make the API RESTful... not implemented again because, example init
	fmt.Println(result)
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("post created!"))
}

func (h *PostsHandler) get(w http.ResponseWriter, r *http.Request) {
	// Not Implemented - would be h.Repo.GetAllPosts() ... you get the idea
	w.Write([]byte("Here are your posts!"))
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/comments
// Allowed methods: GET, POST, PUT, DELETE

type CommentsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewCommentsHandler(r repo.IRepository) *CommentsHandler {
	return &CommentsHandler{Repo: r}
}

// A CommentsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *CommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentsHandler) post(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received comment:", comment.Body)

	// Example Comment to test function
	// comment := models.Comment{
	// 	Body: "Example",
	// 	CreatedAt: 111111,
	//	PostId: 3,
	// 	UpdatedAt: 111111,
	// 	UserId: 2}

	// Validate the post
	if validationErr := comment.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.CreateComment(comment)
	if createErr != nil {
		log.Println("Failed to create comment in the repository:", createErr)
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment created successfully!"))
}

func (h *CommentsHandler) get(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	allPosts, err := h.Repo.GetAllPosts()
	if err != nil {
		log.Println("Failed to get comments in CommentHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your comments"))
}

func (h *CommentsHandler) put(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received post:", comment.Body)

	// Example Comment to test function
	// comment := models.Comment{
	// 	Body: "Updated Example",
	// 	CreatedAt: 111111,
	// 	UpdatedAt: 333333,
	// 	UserId: 2}

	// Validate the post
	if validationErr := comment.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.UpdateComment(comment)
	if createErr != nil {
		log.Println("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment updated successfully!"))
}

func (h *CommentsHandler) delete(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// figure out postId
	var commentId int
	err := json.NewDecoder(r.Body).Decode(&commentId)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received delete request for coomentId:", commentId)

	// example postId for testing
	// postId := 1

	err = h.Repo.DeleteCommentById(commentId)
	if err != nil {
		log.Println("Failed to delete Comment. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("comment was deleted"))
}

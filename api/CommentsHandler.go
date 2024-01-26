package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/comments
// Allowed methods: POST

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

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// HTTP GET logic
	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentsHandler) post(w http.ResponseWriter, r *http.Request) {

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
	// 	Body:      "Example Comment",
	// 	CreatedAt: 1111111,
	// 	PostId:    2,
	// 	UpdatedAt: 1111111,
	// 	UserId:    1}

	// Validate the comment
	if validationErr := comment.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create comment in the repository
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

	allComments, err := h.Repo.GetAllComments()
	if err != nil {
		log.Println("Failed to get comments in CommentHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allComments)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your comments"))
}

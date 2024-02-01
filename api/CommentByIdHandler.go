package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
	"strings"
)

// Endpoint: /api/comments/{commentId}
// Allowed methods: GET, PUT, DELETE
type CommentByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewCommentByIdHandler(r repo.IRepository) *CommentByIdHandler {
	return &CommentByIdHandler{Repo: r}
}

// A CommentsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *CommentByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	// queryParams := r.URL.Query()
	// commentIdString := queryParams.Get("commentId")
	// commentId, postIdErr := strconv.Atoi(commentIdString)
	fields := strings.Split(r.URL.Path, "/")
	commentId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	comment, err := h.Repo.GetCommentById(commentId)
	if err != nil {
		log.Println("Failed to get comments in GetCommentByIdHandler. ", err.Error())
		http.Error(w, "failed to retrieve comment from db", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *CommentByIdHandler) put(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received comment to update:", comment.Body)

	// Validate the comment
	if validationErr := comment.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateComment(comment)
	if createErr != nil {
		log.Println("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *CommentByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	// figure out userID
	queryParams := r.URL.Query()
	// fmt.Println(queryParams)
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		log.Println("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	err := h.Repo.DeleteCommentById(userID)
	if err != nil {
		log.Println("Failed to delete Comments. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

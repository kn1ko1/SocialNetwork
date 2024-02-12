package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/posts/{postId}/comments
// Allowed methods: GET, DELETE

type CommentsByPostIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewCommentsByPostIdHandler(r repo.IRepository) *CommentsByPostIdHandler {
	return &CommentsByPostIdHandler{Repo: r}
}

// A CommentsByPostIdHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *CommentsByPostIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentsByPostIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	postId, err := strconv.Atoi(fields[len(fields)-2])
	if err != nil {
		utils.HandleError("Problem with AtoI postID. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Retrieve comments by post ID from the repository
	comments, err := h.Repo.GetCommentsByPostId(postId)
	if err != nil {
		utils.HandleError("Failed to get comments by post ID. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Encode and write the response
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *CommentsByPostIdHandler) delete(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	postId, err := strconv.Atoi(fields[len(fields)-2])
	if err != nil {
		utils.HandleError("Atoi Error.", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Delete comments by post ID from the repository
	err = h.Repo.DeleteCommentsByPostId(postId)
	if err != nil {
		utils.HandleError("Failed to delete comments by post ID. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Comments for the post were deleted"))
}

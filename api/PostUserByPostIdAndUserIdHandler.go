package api

import (
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
)

// Allowed methods: GET, PUT, DELETE

type PostUserByPostIdAndUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostUserByPostIdAndUserIdHandler(r repo.IRepository) *PostUserByPostIdAndUserIdHandler {
	return &PostUserByPostIdAndUserIdHandler{Repo: r}
}

// A PostssHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostUserByPostIdAndUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// case http.MethodGet:
	// 	h.get(w, r)
	// 	return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func (h *PostUserByPostIdAndUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out userId
	queryParams := r.URL.Query()
	postIdString := queryParams.Get("postId")
	postId, postIdErr := strconv.Atoi(postIdString)
	if postIdErr != nil {
		utils.HandleError("Problem with AtoI userId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userIdString := queryParams.Get("userId")
	userId, userIdErr := strconv.Atoi(userIdString)
	if userIdErr != nil {
		utils.HandleError("Problem with AtoI userId. ", userIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for postId", postId, ", userId", userId)

	err := h.Repo.DeletePostUserByPostIdAndUserId(postId, userId)
	if err != nil {
		utils.HandleError("Failed to delete posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

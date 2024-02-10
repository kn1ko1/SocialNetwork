package api

import (
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
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

	fields := strings.Split(r.URL.Path, "/")
	postIdStr := fields[len(fields)-1]
	userIdStr := fields[len(fields)-3]

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleError("Invalid User ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for postId", postId, ", userId", userId)

	err = h.Repo.DeletePostUserByPostIdAndUserId(postId, userId)
	if err != nil {
		utils.HandleError("Failed to delete posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint:   ?
// Allowed methods: GET, DELETE

type PostUsersByPostIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostUsersByPostIdHandler(r repo.IRepository) *PostUsersByPostIdHandler {
	return &PostUsersByPostIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostUsersByPostIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodPost:
		h.get(w, r)
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

func (h *PostUsersByPostIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	postId, postIdErr := strconv.Atoi(fields[len(fields)-1])
	if postIdErr != nil {
		utils.HandleError("Problem with AtoI postId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	postUsers, err := h.Repo.GetPostUsersByPostId(postId)
	if err != nil {
		utils.HandleError("Failed to get eventUsers in PostUserByPostIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(postUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *PostUsersByPostIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	postId, postIdErr := strconv.Atoi(fields[len(fields)-1])
	if postIdErr != nil {
		utils.HandleError("Problem with AtoI postId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := h.Repo.DeletePostUsersByPostId(postId)
	if err != nil {
		utils.HandleError("Failed to delete postUsers. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

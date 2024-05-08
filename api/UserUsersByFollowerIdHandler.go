package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/event/{postId}/eventUser   ?
// Allowed methods: GET, DELETE

type UserUsersByFollowerIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserUsersByFollowerIdHandler(r repo.IRepository) *UserUsersByFollowerIdHandler {
	return &UserUsersByFollowerIdHandler{Repo: r}
}

// implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserUsersByFollowerIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodGet:
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

func (h *UserUsersByFollowerIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	followerId, followerIdErr := strconv.Atoi(fields[len(fields)-2])
	if followerIdErr != nil {
		utils.HandleError("Problem with AtoI followerId. ", followerIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userUsers, err := h.Repo.GetUserUsersByFollowerId(followerId)
	if err != nil {
		utils.HandleError("Failed to get userUsers in UserUserByFollowerIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *UserUsersByFollowerIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	followerId, followerIdErr := strconv.Atoi(fields[len(fields)-1])
	if followerIdErr != nil {
		utils.HandleError("Problem with AtoI followerId. ", followerIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := h.Repo.DeleteUserUsersByFollowerId(followerId)
	if err != nil {
		utils.HandleError("Failed to delete userUsers. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

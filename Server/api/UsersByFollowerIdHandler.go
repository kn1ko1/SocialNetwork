package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/event/{postId}/eventUser   ?
// Allowed methods: GET, DELETE

type UsersByFollowerIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUsersByFollowerIdHandler(r repo.IRepository) *UsersByFollowerIdHandler {
	return &UsersByFollowerIdHandler{Repo: r}
}

// implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UsersByFollowerIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodGet:
		h.get(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UsersByFollowerIdHandler) get(w http.ResponseWriter, r *http.Request) {
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

	var users []models.User
	for i := 0; i < len(userUsers); i++ {
		user, err := h.Repo.GetUserById(userUsers[i].SubjectId)
		if err != nil {
			utils.HandleError("Failed to get user in GetUserById. ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

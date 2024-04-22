package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/comments/{commentId}
// Allowed methods: GET, PUT, DELETE
type UsersTransportHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUsersTransportHandler(r repo.IRepository) *UsersTransportHandler {
	return &UsersTransportHandler{Repo: r}
}

// A CommentsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UsersTransportHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		// Add auth - must be creator
		h.get(w)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UsersTransportHandler) get(w http.ResponseWriter) {
	allUsers, err := h.Repo.GetAllUsersTransport()
	if err != nil {
		utils.HandleError("Unable to get all users in UsersTransportHandler:", err)
		http.Error(w, "Unable to get all users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

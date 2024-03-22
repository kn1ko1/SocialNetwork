package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Allowed methods: GET, POST

type UsersByPublicHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUsersByPublicHandler(r repo.IRepository) *UsersByPublicHandler {
	return &UsersByPublicHandler{Repo: r}
}

// A UsersByPublicHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UsersByPublicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		h.get(w)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UsersByPublicHandler) get(w http.ResponseWriter) {

	userUsers, err := h.Repo.GetUsersByPublic()
	if err != nil {
		utils.HandleError("Failed to get all Users. ", err)
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

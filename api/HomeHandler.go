package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/home  ?
// Allowed methods: GET

type HomeHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewHomeHandler(r repo.IRepository) *HomeHandler {
	return &HomeHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *HomeHandler) get(w http.ResponseWriter, r *http.Request) {

	user, userErr := getUser(r)
	if userErr != nil {
		utils.HandleError("Problem getting user.", userErr)
	}

	homeData, err := h.Repo.GetHomeDataForUser(user.UserId)
	if err != nil {
		utils.HandleError("Failed to get homeData in HomeHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(homeData)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

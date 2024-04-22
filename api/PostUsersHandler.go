package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/users/posts
// Allowed methods: GET

type PostUsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostUsersHandler(r repo.IRepository) *PostUsersHandler {
	return &PostUsersHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostUsersHandler) post(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	userPosts, err := h.Repo.GetPostsByUserId(user.UserId)
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostsByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userPosts)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

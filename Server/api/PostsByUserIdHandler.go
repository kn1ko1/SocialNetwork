package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/groups/{userId}/posts
// Allowed methods: GET, DELETE

type PostsByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsByUserIdHandler(r repo.IRepository) *PostsByUserIdHandler {
	return &PostsByUserIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userId, err := strconv.Atoi(fields[len(fields)-2])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupPosts, err := h.Repo.GetPostsByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get posts in PostsByUserIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(groupPosts)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

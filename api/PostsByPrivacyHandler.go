package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/posts/{postId}
// Allowed methods: GET

type PostByPrivacyHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostByPrivacyHandler(r repo.IRepository) *PostByPrivacyHandler {
	return &PostByPrivacyHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostByPrivacyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostByPrivacyHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	privacyString := queryParams.Get("privacy")
	userPosts, err := h.Repo.GetPostsByPrivacy(privacyString)
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostByPrivacyHandler. ", err)
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

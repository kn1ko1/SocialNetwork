package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/utils"
	"strings"
)

// Endpoint: /api/posts/privacy/public
// Allowed methods: GET

type PostsByPrivacyHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsByPrivacyHandler(r repo.IRepository) *PostsByPrivacyHandler {
	return &PostsByPrivacyHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsByPrivacyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByPrivacyHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	privacyString := (fields[len(fields)-1])

	if privacyString != "public" && privacyString != "private" {
		utils.HandleError("Failed to get posts in GetPostsByPrivacyHandler", errors.New("invalid privacy string"))
		http.Error(w, "invalid privacy string", http.StatusBadRequest)
	}

	userPosts, err := h.Repo.GetPostsByPrivacy(privacyString)
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostsByPrivacyHandler. ", err)
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

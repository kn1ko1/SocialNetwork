package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
)

// Endpoint: /api/users/posts
// Allowed methods: GET

type UserPostsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserPostsHandler(r repo.IRepository) *UserPostsHandler {
	return &UserPostsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserPostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserPostsHandler) get(w http.ResponseWriter, r *http.Request) {
	// queryParams := r.URL.Query()
	// postIdString := queryParams.Get("postId")
	// postId, postIdErr := strconv.Atoi(postIdString)
	user, err := getUser(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userPosts, err := h.Repo.GetPostsByUserId(user.UserId)
	if err != nil {
		log.Println("Failed to get posts in GetPostsByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// LEAVE UNIMPLEMENTED IN SWITCH CASE
// func (h *UserPostsHandler) delete(w http.ResponseWriter, r *http.Request) {
// 	user, err := getUser(r)
// 	if err != nil {
// 		http.Error(w, "unauthorized", http.StatusUnauthorized)
// 		return
// 	}
// 	// log.Println("Received delete request for userID:", userID)
// 	err = h.Repo.DeletePostsByUserId(user.UserId)
// 	if err != nil {
// 		log.Println("Failed to delete Posts. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

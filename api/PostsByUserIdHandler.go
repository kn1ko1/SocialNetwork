package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: /api/users/{userId}/posts
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
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	postIdString := queryParams.Get("postId")
	postId, postIdErr := strconv.Atoi(postIdString)
	if postIdErr != nil {
		log.Println("Problem with AtoI postId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userPosts, err := h.Repo.GetPostById(postId)
	if err != nil {
		log.Println("Failed to get posts in GetPostsByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
}

func (h *PostsByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		log.Println("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	err := h.Repo.DeletePostsByUserId(userID)
	if err != nil {
		log.Println("Failed to delete Posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/repo"
)

// Endpoint: /api/users/posts
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
	// queryParams := r.URL.Query()
	// postIdString := queryParams.Get("postId")
	// postId, postIdErr := strconv.Atoi(postIdString)
	c, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := auth.AuthenticateSessionCookie(c)
	if err != nil {
		log.Println(err.Error())
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

func (h *PostsByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := auth.AuthenticateSessionCookie(c)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// log.Println("Received delete request for userID:", userID)
	err = h.Repo.DeletePostsByUserId(user.UserId)
	if err != nil {
		log.Println("Failed to delete Posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

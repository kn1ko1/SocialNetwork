package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/posts
// GET /api/posts - Retrieves all posts from DB
// POST /api/posts - Create a new post

// Allowed methods: GET, POST

type PostsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsHandler(r repo.IRepository) *PostsHandler {
	return &PostsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsHandler) post(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received post:", post.UserId, post.Body)

	// Validate the post
	if validationErr := post.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.CreatePost(post)
	if createErr != nil {
		log.Println("Failed to create post in the repository:", createErr)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post created successfully!"))
}

func (h *PostsHandler) get(w http.ResponseWriter, r *http.Request) {

	allPosts, err := h.Repo.GetAllPosts()
	if err != nil {
		log.Println("Failed to get posts in PostHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

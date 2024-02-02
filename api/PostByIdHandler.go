package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/posts/{postId}
// Allowed methods: GET, PUT, DELETE

type PostByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostByIdHandler(r repo.IRepository) *PostByIdHandler {
	return &PostByIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	postId, postIdErr := strconv.Atoi(fields[len(fields)-1])
	if postIdErr != nil {
		utils.HandleError("Problem with AtoI postId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userPosts, err := h.Repo.GetPostById(postId)
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostByIdHandler. ", err)
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

func (h *PostByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating:", post.PostId, post.Body)

	// Validate the User <3
	if validationErr := post.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update user in the repository
	result, createErr := h.Repo.UpdatePost(post)
	if createErr != nil {
		utils.HandleError("Failed to update user in the repository:", createErr)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *PostByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		utils.HandleError("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	err := h.Repo.DeletePostById(userID)
	if err != nil {
		utils.HandleError("Failed to delete Posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

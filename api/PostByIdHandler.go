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
	postIdStr := fields[len(fields)-1]

	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		utils.HandleError("Invalid post ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received get request for post Id:", postId)
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

	parseMultipartFormErr := r.ParseMultipartForm(10 << 20)
	if parseMultipartFormErr != nil {
		utils.HandleError("Unable to Parse Multipart Form.", parseMultipartFormErr)
	}

	file, fileHeader, formFileErr := r.FormFile("image")
	if formFileErr != nil {
		utils.HandleError("Error reading image.", formFileErr)
	}

	//if file is given
	if file != nil {
		defer file.Close()
		var ImageProcessingrErr error
		post.ImageURL, ImageProcessingrErr = ImageProcessing(w, r, file, *fileHeader)
		if ImageProcessingrErr != nil {
			utils.HandleError("Error with ImageHandler", ImageProcessingrErr)
		}
	}

	// Update user in the repository
	result, createErr := h.Repo.UpdatePost(post)
	if createErr != nil {
		utils.HandleError("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
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
	fields := strings.Split(r.URL.Path, "/")
	postIdStr := fields[len(fields)-1]

	postId, postIdErr := strconv.Atoi(postIdStr)
	if postIdErr != nil {
		utils.HandleError("Problem with AtoI userID. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", postId)

	err := h.Repo.DeletePostById(postId)
	if err != nil {
		utils.HandleError("Failed to delete Posts. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

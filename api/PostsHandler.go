package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/posts
// Allowed methods: POST

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
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
		// case http.MethodGet:
		// 	h.get(w, r)
		// 	return
		// default:
		// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		// 	return
	}
}

func (h *PostsHandler) post(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received post:", post.UserId, post.Body)

	// Validate the post
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
		fmt.Println("POST INSERTED WITH FILE")
	} else {
		fmt.Println("POST INSERTED WITHOUT FILE")
	}

	result, createErr := h.Repo.CreatePost(post)
	if createErr != nil {
		utils.HandleError("Failed to create post in the repository:", createErr)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *PostsHandler) get(w http.ResponseWriter, r *http.Request) {

// 	allPosts, err := h.Repo.GetAllPosts()
// 	if err != nil {
// 		utils.HandleError("Failed to get posts in PostHandler. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(allPosts)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

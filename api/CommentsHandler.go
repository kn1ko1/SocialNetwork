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

// Endpoint: /api/comments
// Allowed methods: POST

type CommentsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewCommentsHandler(r repo.IRepository) *CommentsHandler {
	return &CommentsHandler{Repo: r}
}

// A CommentsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *CommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "unauthorized", http.StatusUnauthorized)
	// Switch on the Request method, call the correct subroutine...
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

func (h *CommentsHandler) post(w http.ResponseWriter, r *http.Request) {

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received comment:", comment.Body)

	// Validate the comment
	if validationErr := comment.Validate(); validationErr != nil {
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

	defer file.Close()

	//if file is given
	if file != nil {
		var imageHandlerErr error
		comment.ImageURL, imageHandlerErr = ImageProcessing(w, r, file, *fileHeader)
		if imageHandlerErr != nil {
			utils.HandleError("Error with ImageHandler", imageHandlerErr)
		}
		fmt.Println("COMMENT INSERTED WITH FILE")
	} else {
		fmt.Println("COMMENT INSERTED WITHOUT FILE")
	}

	// Create comment in the repository
	result, createErr := h.Repo.CreateComment(comment)
	if createErr != nil {
		utils.HandleError("Failed to create comment in the repository:", createErr)
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Comment created successfully!"))
}

// func (h *CommentsHandler) get(w http.ResponseWriter, r *http.Request) {

// 	allComments, err := h.Repo.GetAllComments()
// 	if err != nil {
// 		utils.HandleError("Failed to get comments in CommentHandler. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(allComments)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// }

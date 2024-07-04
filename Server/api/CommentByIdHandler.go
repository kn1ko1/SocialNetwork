package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/imageProcessing"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/comments/{commentId}
// Allowed methods: GET, PUT, DELETE
type CommentByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewCommentByIdHandler(r repo.IRepository) *CommentByIdHandler {
	return &CommentByIdHandler{Repo: r}
}

// A CommentsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *CommentByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		// Add auth - must be creator
		h.get(w, r)
		return
	case http.MethodDelete:
		// Add auth - user must be creator
		h.delete(w, r)
		return
	case http.MethodPut:
		// Add auth - user must be creator
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	commentId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		utils.HandleError("Atoi Error.", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	comment, err := h.Repo.GetCommentById(commentId)
	if err != nil {
		utils.HandleError("Failed to get comments in GetCommentByIdHandler. ", err)
		http.Error(w, "failed to retrieve comment from db", http.StatusInternalServerError)
		return
	}
	// if comment.UserId != user.Id {
	//		unauthorized error
	//}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *CommentByIdHandler) put(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received comment to update:", comment.Body)

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

	//if file is given
	if file != nil {
		defer file.Close()
		var ImageProcessingrErr error
		comment.ImageURL, ImageProcessingrErr = imageProcessing.ImageProcessing(w, r, file, *fileHeader)
		if ImageProcessingrErr != nil {
			utils.HandleError("Error with ImageHandler", ImageProcessingrErr)
		}
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateComment(comment)
	if createErr != nil {
		utils.HandleError("Failed to update comment in the repository:", createErr)
		http.Error(w, "Failed to update comment", http.StatusInternalServerError)
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

func (h *CommentByIdHandler) delete(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	commentId, commentIdErr := strconv.Atoi(fields[len(fields)-1])
	if commentIdErr != nil {
		utils.HandleError("Problem with AtoI commentId. ", commentIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for commentId:", commentId)

	err := h.Repo.DeleteCommentById(commentId)
	if err != nil {
		utils.HandleError("Failed to delete Comments. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

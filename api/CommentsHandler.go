package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
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
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *CommentsHandler) post(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie(auth.CookieName)
	if err != nil {

		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}

	user, exists := auth.SessionMap[cookie.Value]
	if !exists {
		utils.HandleError("Error finding User, need to log in again", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	contentType := r.Header.Get("Content-Type")
	var comment models.Comment
	switch contentType {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			utils.HandleError("Failed to decode request body:", err)
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			utils.HandleError("Failed to parse form:", err)
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}
		comment.Body = r.PostFormValue("comment-body")
		ctime := time.Now().UTC().UnixMilli()
		comment.CreatedAt = ctime
		comment.UpdatedAt = ctime
		comment.ImageURL = ""
		comment.PostId = 1
		comment.UserId = user.UserId
	}
	// Validate the post
	if validationErr := comment.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}
	log.Println("Received post:", comment.UserId, comment.Body)
	// parseMultipartFormErr := r.ParseMultipartForm(10 << 20)
	// if parseMultipartFormErr != nil {
	// 	utils.HandleError("Unable to Parse Multipart Form.", parseMultipartFormErr)
	// }
	// file, fileHeader, formFileErr := r.FormFile("image")
	// if formFileErr != nil {
	// 	utils.HandleError("Error reading image.", formFileErr)
	// }

	// //if file is given
	// if file != nil {
	// 	defer file.Close()
	// 	var ImageProcessingrErr error
	// 	comment.ImageURL, ImageProcessingrErr = ImageProcessing(w, r, file, *fileHeader)
	// 	if ImageProcessingrErr != nil {
	// 		utils.HandleError("Error with ImageHandler", ImageProcessingrErr)
	// 	}
	// 	fmt.Println("POST INSERTED WITH FILE")
	// } else {
	// 	fmt.Println("POST INSERTED WITHOUT FILE")
	// }

	result, createErr := h.Repo.CreateComment(comment)
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

func (h *CommentsHandler) get(w http.ResponseWriter) {

	allComments, err := h.Repo.GetAllComments()
	if err != nil {
		utils.HandleError("Failed to get comments in CommentHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allComments)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

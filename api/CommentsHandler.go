package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
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
	ctime := time.Now().UTC().UnixMilli()

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

	// Parse form data
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.HandleError("Failed to parse form data:", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	body := r.FormValue("body")
	imageURL := ""
	postId, _ := strconv.Atoi(r.FormValue("postId"))

	userId := user.UserId

	log.Println("[api/CommentsHandler] Received Post:", body)
	log.Println("[api/CommentsHandler] PostId:", postId)
	log.Println("[api/CommentsHandler] UserId:", userId)

	// Handle file upload
	file, fileHeader, _ := r.FormFile("image")
	if file != nil {

		defer file.Close()
		imageURL, err = ImageProcessing(w, r, file, *fileHeader)
		if err != nil {
			utils.HandleError("Error with ImageHandler", err)
			// http.Error(w, "Failed to process image", http.StatusInternalServerError)
			return
		}
		log.Println("[api/PostsHandler] Image Stored at:", imageURL)
	}
	comment := models.Comment{
		Body:      body,
		CreatedAt: ctime,
		ImageURL:  imageURL,
		PostId:    postId,
		UpdatedAt: ctime,
		UserId:    userId,
	}
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

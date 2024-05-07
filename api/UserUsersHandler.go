package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
)

// Allowed methods: POST

type UserUsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserUsersHandler(r repo.IRepository) *UserUsersHandler {
	return &UserUsersHandler{Repo: r}
}

// A UserUsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserUsersHandler) post(w http.ResponseWriter, r *http.Request) {
	var userUser models.UserUser

	ctime := time.Now().UTC().UnixMilli()
	// cookie, err := r.Cookie(auth.CookieName)
	// if err != nil {

	// 	utils.HandleError("Error verifying cookie", err)
	// 	http.Redirect(w, r, "auth/login", http.StatusSeeOther)
	// 	return
	// }

	// follower, exists := auth.SessionMap[cookie.Value]
	// if !exists {
	// 	utils.HandleError("Error finding User, need to log in again", err)
	// 	http.Redirect(w, r, "auth/login", http.StatusSeeOther)
	// 	return
	// }

	// decodes subjectId directly into userUser struct
	if err := json.NewDecoder(r.Body).Decode(&userUser); err != nil {
		utils.HandleError("unable to decode subjectId. ", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	userUser.CreatedAt = ctime
	userUser.UpdatedAt = ctime

	log.Println("[api/UserUsersHandler] Following.  FollowerId:", userUser.FollowerId, ". SubjectId:", userUser.SubjectId)

	// Validate the post
	if validationErr := userUser.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	result, createErr := h.Repo.CreateUserUser(userUser)
	if createErr != nil {
		utils.HandleError("Failed to create post in the repository:", createErr)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *UserUsersHandler) get(w http.ResponseWriter, r *http.Request) {

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

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
	"strings"
	"time"
)

// Allowed methods: GET, PUT, DELETE

type UserUserBySubjectIdAndFollowerIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserUserBySubjectIdAndFollowerIdHandler(r repo.IRepository) *UserUserBySubjectIdAndFollowerIdHandler {
	return &UserUserBySubjectIdAndFollowerIdHandler{Repo: r}
}

// implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserUserBySubjectIdAndFollowerIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	// case http.MethodGet:
	// 	h.get(w, r)
	// 	return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserUserBySubjectIdAndFollowerIdHandler) post(w http.ResponseWriter, r *http.Request) {
	ctime := time.Now().UTC().UnixMilli()
	cookie, err := r.Cookie(auth.CookieName)
	if err != nil {

		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}

	follower, exists := auth.SessionMap[cookie.Value]
	if !exists {
		utils.HandleError("Error finding User, need to log in again", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	fields := strings.Split(r.URL.Path, "/")
	subjectIdStr := fields[len(fields)-1]

	subjectId, err := strconv.Atoi(subjectIdStr)
	if err != nil {
		utils.HandleError("Invalid subject ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	userUser := models.UserUser{
		CreatedAt:  ctime,
		FollowerId: follower.UserId,
		SubjectId:  subjectId,
		UpdatedAt:  ctime,
	}

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
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *UserUserBySubjectIdAndFollowerIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	subjectIdStr := fields[len(fields)-1]

	subjectId, err := strconv.Atoi(subjectIdStr)
	if err != nil {
		utils.HandleError("Invalid subject ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	follower, err := getUser(r)
	followerId := follower.UserId
	if err != nil {
		utils.HandleError("Failed to get user. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for subjectId", subjectId, ", followerId", followerId)

	err = h.Repo.DeleteUserUserBySubjectIdAndFollowerId(subjectId, followerId)
	if err != nil {
		utils.HandleError("Failed to delete useruser. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user user was deleted"))
}

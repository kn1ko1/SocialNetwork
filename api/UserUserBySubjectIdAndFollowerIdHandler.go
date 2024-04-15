package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
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
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserUserBySubjectIdAndFollowerIdHandler) get(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	subjectIdStr := fields[len(fields)-1]
	subjectId, err := strconv.Atoi(subjectIdStr)
	if err != nil {
		utils.HandleError("Invalid subject ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	followerIdStr := fields[len(fields)-3]
	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		utils.HandleError("Invalid follower ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	result, getErr := h.Repo.GetUserUserByFollowerIdAndSubjectId(followerId, subjectId)
	if getErr != nil {
		// If no follower is found in the database then returns a 404 error
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			utils.HandleError("Failed to encode and write JSON response. ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		return
	}
	log.Println("[api/UserUsersHandler] Found Follow.  FollowerId:", followerId, ". SubjectId:", subjectId)

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

	followerIdStr := fields[len(fields)-3]
	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		utils.HandleError("Invalid follower ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	err = h.Repo.DeleteUserUserBySubjectIdAndFollowerId(subjectId, followerId)
	if err != nil {
		utils.HandleError("Failed to delete useruser. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("[api/UserUserBySubjectIdAndFollowerIdHandler] Unfollowing.  FollowerId:", followerId, ". SubjectId:", subjectId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user user was deleted"))
}

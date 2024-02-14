package api

import (
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

func (h *UserUserBySubjectIdAndFollowerIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	subjectIdStr := fields[len(fields)-1]
	followerIdStr := fields[len(fields)-3]

	subjectId, err := strconv.Atoi(subjectIdStr)
	if err != nil {
		utils.HandleError("Invalid subject ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	followerId, err := strconv.Atoi(followerIdStr)
	if err != nil {
		utils.HandleError("Invalid follower ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
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

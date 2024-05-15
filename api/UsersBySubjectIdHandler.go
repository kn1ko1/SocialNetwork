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

// Endpoint: /api/event/{postId}/eventUser   ?
// Allowed methods: GET, DELETE

type UsersBySubjectIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUsersBySubjectIdHandler(r repo.IRepository) *UsersBySubjectIdHandler {
	return &UsersBySubjectIdHandler{Repo: r}
}

// implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UsersBySubjectIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodGet:
		h.get(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UsersBySubjectIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	subjectId, subjectIdErr := strconv.Atoi(fields[len(fields)-2])
	if subjectIdErr != nil {
		utils.HandleError("Problem with AtoI subjectId. ", subjectIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userUsers, err := h.Repo.GetUserUsersBySubjectId(subjectId)
	if err != nil {
		utils.HandleError("Failed to get userUsers in UserUserBysubjectIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("followed", userUsers)
	var users []models.User
	for i := 0; i < len(userUsers); i++ {
		user, err := h.Repo.GetUserById(userUsers[i].FollowerId)
		if err != nil {
			utils.HandleError("Failed to get user in GetUserById. ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		users = append(users, user)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

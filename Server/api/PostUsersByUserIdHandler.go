package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/postsUsers/users/{userId}   ?
// Allowed methods: GET, DELETE

type PostUsersByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostUsersByUserIdHandler(r repo.IRepository) *PostUsersByUserIdHandler {
	return &PostUsersByUserIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostUsersByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodPost:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostUsersByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userIdStr := fields[len(fields)-1]

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	eventUsers, err := h.Repo.GetPostUsersByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get eventUsers in PostUserByUserIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(eventUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *PostUsersByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	userIdString := queryParams.Get("userId")
	userId, userIdErr := strconv.Atoi(userIdString)
	if userIdErr != nil {
		utils.HandleError("Problem with AtoI messageId. ", userIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for messageId:", userId)

	err := h.Repo.DeletePostUsersByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to delete eventUsers. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"strconv"
)

// Allowed methods: GET, DELETE

type GroupUserByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUserByIdHandler(r repo.IRepository) *GroupUserByIdHandler {
	return &GroupUserByIdHandler{Repo: r}
}

// A GroupUsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupUserByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *GroupUserByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	GroupUserIdString := queryParams.Get("GroupUserId")
	GroupUserId, GroupUserIdErr := strconv.Atoi(GroupUserIdString)
	if GroupUserIdErr != nil {
		log.Println("Problem with AtoI GroupUserId. ", GroupUserIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	GroupUser, err := h.Repo.GetGroupUser(GroupUserId)
	if err != nil {
		log.Println("Failed to get GroupUser in GetGroupUserByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(GroupUser)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GroupUserByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	GroupUserIdString := queryParams.Get("GroupUserId")
	GroupUserId, GroupUserIdErr := strconv.Atoi(GroupUserIdString)
	if GroupUserIdErr != nil {
		log.Println("Problem with AtoI GroupUserId. ", GroupUserIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for GroupUserId:", GroupUserId)

	err := h.Repo.DeleteGroupUser(GroupUserId)
	if err != nil {
		log.Println("Failed to delete GroupUsers. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

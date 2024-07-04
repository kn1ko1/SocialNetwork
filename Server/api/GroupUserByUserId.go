package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, DELETE

type GroupUsersByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUsersByUserIdHandler(r repo.IRepository) *GroupUsersByUserIdHandler {
	return &GroupUsersByUserIdHandler{Repo: r}
}

// A GroupUsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupUsersByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	// case http.MethodDelete:
	// 	h.delete(w, r)
	// 	return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GroupUsersByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userIdStr := fields[len(fields)-2]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleError("Invalid GroupUser ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	GroupUsers, err := h.Repo.GetGroupUsersByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get GroupUser in GetGroupUsersByUserIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(GroupUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *GroupUsersByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	fields := strings.Split(r.URL.Path, "/")
// 	groupIdStr := fields[len(fields)-2]

// 	groupUserId, err := strconv.Atoi(groupIdStr)
// 	if err != nil {
// 		utils.HandleError("Invalid Group ID. ", err)
// 		http.Error(w, "internal server errror", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Received delete request for GroupUserId:", groupUserId)

// 	err = h.Repo.DeleteGroupUsersByGroupId(groupId)
// 	if err != nil {
// 		utils.HandleError("Failed to delete GroupUsers. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

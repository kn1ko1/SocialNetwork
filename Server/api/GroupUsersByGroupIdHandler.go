package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, DELETE

type GroupUsersByGroupIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUsersByGroupIdHandler(r repo.IRepository) *GroupUsersByGroupIdHandler {
	return &GroupUsersByGroupIdHandler{Repo: r}
}

// A GroupUsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupUsersByGroupIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *GroupUsersByGroupIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	groupdStr := fields[len(fields)-2]
	groupId, err := strconv.Atoi(groupdStr)
	if err != nil {
		utils.HandleError("Invalid GroupUser ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	GroupUsers, err := h.Repo.GetGroupUsersByGroupId(groupId)
	if err != nil {
		utils.HandleError("Failed to get GroupUser in GetGroupUsersByGroupIdHandler. ", err)
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

// func (h *GroupUsersByGroupIdHandler) delete(w http.ResponseWriter, r *http.Request) {

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

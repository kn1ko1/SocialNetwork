package api

import (
	"log"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, PUT, DELETE

type GroupUserByGroupIdAndUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUserByGroupIdAndUserIdHandler(r repo.IRepository) *GroupUserByGroupIdAndUserIdHandler {
	return &GroupUserByGroupIdAndUserIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupUserByGroupIdAndUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
func (h *GroupUserByGroupIdAndUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	groupIdStr := fields[len(fields)-4]
	userIdStr := fields[len(fields)-1]

	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for eventId", groupId, ", userId", userId)

	err = h.Repo.DeleteGroupUserByGroupIdAndUserId(groupId, userId)
	if err != nil {
		utils.HandleError("Failed to delete Events. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

// func (h *GroupUserByGroupIdAndUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
// 	fields := strings.Split(r.URL.Path, "/")
// 	groupIdStr := fields[len(fields)-2]
// 	userIdStr := fields[len(fields)-2]

// 	groupId, err := strconv.Atoi(groupIdStr)
// 	if err != nil {
// 		utils.HandleError("Invalid Group ID. ", err)
// 		http.Error(w, "internal server errror", http.StatusInternalServerError)
// 		return
// 	}

// 	userId, err := strconv.Atoi(userIdStr)
// 	if err != nil {
// 		utils.HandleError("Invalid Group ID. ", err)
// 		http.Error(w, "internal server errror", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Received get request for eventId", groupId, ", userId", userId)

// 	eventUser, err := h.Repo.GetGroupUserByGroupIdAndUserId(groupId, userId)
// 	if err != nil {
// 		utils.HandleError("Failed to get posts in GetEventsByGroupId. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(eventUser)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Here are your eventUsers"))
// }

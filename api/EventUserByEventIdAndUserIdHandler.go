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

type EventUsersByEventIdAnduserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventUserByEventIdAndUserIdHandler(r repo.IRepository) *EventUsersByEventIdAnduserIdHandler {
	return &EventUsersByEventIdAnduserIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventUsersByEventIdAnduserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
func (h *EventUsersByEventIdAnduserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	eventIdStr := fields[len(fields)-4]
	userIdStr := fields[len(fields)-1]

	eventId, err := strconv.Atoi(eventIdStr)
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
	log.Println("Received delete request for eventId", eventId, ", userId", userId)

	err = h.Repo.DeleteEventUserByEventIdAndUserId(eventId, userId)
	if err != nil {
		utils.HandleError("Failed to delete Events. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

// func (h *EventUsersByEventIdAnduserIdHandler) get(w http.ResponseWriter, r *http.Request) {
// 	fields := strings.Split(r.URL.Path, "/")
// 	eventIdStr := fields[len(fields)-3]
// 	userIdStr := fields[len(fields)-1]

// 	eventId, err := strconv.Atoi(eventIdStr)
// 	if err != nil {
// 		utils.HandleError("Invalid eventId. ", err)
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	userId, err := strconv.Atoi(userIdStr)
// 	if err != nil {
// 		utils.HandleError("Invalid userId. ", err)
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	log.Println("Received delete request for eventId", eventId, ", userId", userId)

// 	eventUser, err := h.Repo.GetEventUserByEventIdanduserId(eventId, userId)
// 	if err != nil {
// 		utils.HandleError("Failed to get posts in GetEventsByGroupId. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(eventUser)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Here are your posts"))
// }

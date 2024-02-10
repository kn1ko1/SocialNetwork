package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/groups/{id}/events
// Allowed methods: GET, PUT, DELETE

type EventsByGroupIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventsByGroupIdHandler(r repo.IRepository) *EventsByGroupIdHandler {
	return &EventsByGroupIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventsByGroupIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventsByGroupIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	groupIdStr := fields[len(fields)-2]
	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	events, err := h.Repo.GetEventsByGroupId(groupId)
	if err != nil {
		utils.HandleError("Failed to get posts in GetEventsByGroupId. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *EventsByGroupIdHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	// look at penultimate id for userId

// 	fields := strings.Split(r.URL.Path, "/")
// 	groupId, userIdErr := strconv.Atoi(fields[len(fields)-1])
// 	if userIdErr != nil {
// 		utils.HandleError("Problem with AtoI groupId. ", userIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Received delete request for groupId:", groupId)

// 	// example postId for testing
// 	// postId := 1

// 	err := h.Repo.DeleteEventsByGroupId(groupId)
// 	if err != nil {
// 		utils.HandleError("Failed to delete Events. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("posts were deleted"))
// }

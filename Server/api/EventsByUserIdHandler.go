package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/groups/{id}/events
// Allowed methods: GET, PUT, DELETE

type EventsByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventsByUserIdHandler(r repo.IRepository) *EventsByUserIdHandler {
	return &EventsByUserIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventsByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *EventsByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userIdStr := fields[len(fields)-2]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleError("Invalid user ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	events, err := h.Repo.GetEventsByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get events in GetEventsByUserId. ", err)
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

func (h *EventsByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	fields := strings.Split(r.URL.Path, "/")
	userId, userIdErr := strconv.Atoi(fields[len(fields)-2])
	if userIdErr != nil {
		utils.HandleError("Problem with AtoI userId. ", userIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userId:", userId)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeleteEventsByUserId(userId)
	if err != nil {
		utils.HandleError("Faidled to delete Events. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("events were deleted"))
}

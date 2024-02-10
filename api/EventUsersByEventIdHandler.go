package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/event/{eventId}/eventUser   ?
// Allowed methods: GET, DELETE

type EventUsersByEventIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventUsersByEventIdHandler(r repo.IRepository) *EventUsersByEventIdHandler {
	return &EventUsersByEventIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventUsersByEventIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodGet:
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

func (h *EventUsersByEventIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	eventIdStr := fields[len(fields)-2]
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		utils.HandleError("Invalid eventId. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	eventUsers, err := h.Repo.GetEventUsersByEventId(eventId)
	if err != nil {
		utils.HandleError("Failed to get eventUsers in EventUserByEventIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(eventUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *EventUsersByEventIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	eventIdStr := fields[len(fields)-2]
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		utils.HandleError("Invalid eventID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for messageId:", eventId)

	err = h.Repo.DeleteEventUsersByEventId(eventId)
	if err != nil {
		utils.HandleError("Failed to delete eventUsers. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

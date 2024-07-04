package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/events/{eventId}
// Allowed methods: GET, PUT, DELETE

type EventByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventByIdHandler(r repo.IRepository) *EventByIdHandler {
	return &EventByIdHandler{Repo: r}
}

func (h *EventByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Add auth - must be a member of the event's group
		h.get(w, r)
		return
	case http.MethodPut:
		// Add auth - must be event creator
		h.put(w, r)
		return
	case http.MethodDelete:
		// Add auth must be event creator
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	eventId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		utils.HandleError("Problem with AtoI eventId. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	event, err := h.Repo.GetEventById(eventId)
	if err != nil {
		utils.HandleError("Failed to get event. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *EventByIdHandler) put(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating event:", event.Title, event.Description)
	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}
	// Create event in the repository
	result, createErr := h.Repo.UpdateEvent(event)
	if createErr != nil {
		utils.HandleError("Failed to update event in the repository:", createErr)
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}
	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *EventByIdHandler) delete(w http.ResponseWriter, r *http.Request) {
	// auth stuff
	fields := strings.Split(r.URL.Path, "/")
	eventId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		utils.HandleError("Problem with AtoI eventId. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = h.Repo.DeleteEventById(eventId)
	if err != nil {
		utils.HandleError("Failed to delete event. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

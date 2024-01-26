package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
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
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	eventIdString := queryParams.Get("eventId")
	eventId, postIdErr := strconv.Atoi(eventIdString)
	if postIdErr != nil {
		log.Println("Problem with AtoI eventId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	event, err := h.Repo.GetEventById(eventId)
	if err != nil {
		log.Println("Failed to get posts in GetPostByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h *EventByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating event:", event.Title, event.Description)

	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.UpdateEventById(event)
	if createErr != nil {
		log.Println("Failed to update event in the repository:", createErr)
		http.Error(w, "Failed to update event", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

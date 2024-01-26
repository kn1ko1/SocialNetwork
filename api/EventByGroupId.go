package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: /api/Messa s/Event/{EventId}
// Allowed methods: GET, PUT, DELETE

type EventByGroupIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventByGroupIdHandler(r repo.IRepository) *EventByGroupIdHandler {
	return &EventByGroupIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventByGroupIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventByGroupIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	EventIdString := queryParams.Get("postId")
	EventId, EventIdErr := strconv.Atoi(EventIdString)
	if EventIdErr != nil {
		log.Println("Problem with AtoI EventId. ", EventIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userEvents, err := h.Repo.GetEventById(EventId)
	if err != nil {
		log.Println("Failed to get posts in GetEventByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userEvents)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
}

func (h *EventByGroupIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var Event models.Event
	err := json.NewDecoder(r.Body).Decode(&Event)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Event to update:", Event.Title)

	// Validate the Event
	if validationErr := Event.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateEvent(Event)
	if createErr != nil {
		log.Println("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Event updated successfully!"))
}

func (h *EventByGroupIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		log.Println("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeleteEventById(userID)
	if err != nil {
		log.Println("Failed to delete Events. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/events
// Allowed methods: GET, POST

type EventsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventsHandler(r repo.IRepository) *EventsHandler {
	return &EventsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return

	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventsHandler) post(w http.ResponseWriter, r *http.Request) {

	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received event:", event.Title, event.Description)

	// Example event to test function
	// event := models.Event{
	// 	CreatedAt:   111111,
	// 	DateTime:    1212121212,
	// 	Description: "example event description",
	// 	GroupID:     1,
	// 	UpdatedAt:   111111,
	// 	Title:       "Magnificient Example Event",
	// 	UserId:      2}

	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.CreateEvent(event)
	if createErr != nil {
		log.Println("Failed to create event in the repository:", createErr)
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
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
	w.Write([]byte("Event created successfully!"))
}

func (h *EventsHandler) get(w http.ResponseWriter, r *http.Request) {

	allEvents, err := h.Repo.GetAllEvents()
	if err != nil {
		log.Println("Failed to get event in EventHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allEvents)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your events"))
}

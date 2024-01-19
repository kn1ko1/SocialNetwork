package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/events
// Allowed methods: GET, POST, PUT, DELETE

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
	// Get Session Cookie
	c, err := r.Cookie("Session")
	if err != nil {
		// Log Error
		log.Println(err.Error())
		// Return HTTP Status Unauthorized
		//
		// N.B. for simplicity of the example, we are simply returning
		// an HTTP error. In the actual project, probably a JSON payload.
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Authenticate Session Cookie - user variable discarded because user struct not used here...
	_, err = auth.AuthenticateSessionCookie(c)
	if err != nil {
		// Same error as above - maker of request is unauthorized
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
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

func (h *EventsHandler) post(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

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

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	allPosts, err := h.Repo.GetAllPosts()
	if err != nil {
		log.Println("Failed to get event in EventHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your events"))
}

func (h *EventsHandler) put(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

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
	// 	Description: "updated example event description",
	// 	GroupID:     1,
	// 	UpdatedAt:   33333333,
	// 	Title:       "Magnificient Updated Example Event",
	// 	UserId:      2}

	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.UpdateEvent(event)
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
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post updated successfully!"))
}

func (h *EventsHandler) delete(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// figure out eventId
	var eventId int
	err := json.NewDecoder(r.Body).Decode(&eventId)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received delete request for eventId:", eventId)

	// example eventId for testing
	// eventId := 1

	err = h.Repo.DeleteEventById(eventId)
	if err != nil {
		log.Println("Failed to delete event. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("event was deleted"))
}

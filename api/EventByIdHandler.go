package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: /api/events/event/{eventId}
// Allowed methods: GET, PUT, DELETE

type EventByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventByUserIdHandler(r repo.IRepository) *EventByUserIdHandler {
	return &EventByUserIdHandler{Repo: r}
}

func (h *EventByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	// case http.MethodDelete:
	// 	h.delete(w, r)
	// 	return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	eventIdString := queryParams.Get("eventId")
	eventId, postIdErr := strconv.Atoi(eventIdString)
	if postIdErr != nil {
		log.Println("Problem with AtoI postId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userPosts, err := h.Repo.GetEventById(eventId)
	if err != nil {
		log.Println("Failed to get posts in GetPostByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
}

func (h *EventByUserIdHandler) put(w http.ResponseWriter, r *http.Request) {

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
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post updated successfully!"))
}

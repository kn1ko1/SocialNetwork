package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/event/{eventId}/eventUser   ?
// Allowed methods: GET, POST

type EventUsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventUsersHandler(r repo.IRepository) *EventUsersHandler {
	return &EventUsersHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return

	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventUsersHandler) post(w http.ResponseWriter, r *http.Request) {

	var eventUser models.EventUser
	err := json.NewDecoder(r.Body).Decode(&eventUser)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received eventUser: User -", eventUser.UserId, "for event -", eventUser.EventId)

	// Validate the eventUser
	if validationErr := eventUser.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create eventUser in the repository
	result, createErr := h.Repo.CreateEventUser(eventUser)
	if createErr != nil {
		utils.HandleError("Failed to create eventUser in the repository:", createErr)
		http.Error(w, "Failed to create eventUser", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

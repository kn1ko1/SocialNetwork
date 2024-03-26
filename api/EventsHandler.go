package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
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
		h.get(w)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventsHandler) post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var event models.Event
	switch contentType {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			utils.HandleError("Failed to decode request body:", err)
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			utils.HandleError("Failed to parse form:", err)
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}
		ctime := time.Now().UTC().UnixMilli()
		event.CreatedAt = ctime
		event.UpdatedAt = ctime
		event.Description = r.PostFormValue("event-description")
		event.Title = r.PostFormValue("event-title")
		event.UserId = 1
		event.GroupId = 1
		t := fmt.Sprintf("%s%s", r.PostFormValue("event-date-time"), ":00Z")
		dtime, err := time.Parse(time.RFC3339, t)
		if err != nil {
			utils.HandleError("Failed to parse date-time data", err)
			http.Error(w, "Failed to parse date-time", http.StatusInternalServerError)
			return
		}
		event.DateTime = dtime.UTC().UnixMilli()
	}
	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}
	log.Println("Received event:", event.Title, event.Description)
	// Create event in the repository
	result, createErr := h.Repo.CreateEvent(event)
	if createErr != nil {
		utils.HandleError("Failed to create event in the repository:", createErr)
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *EventsHandler) get(w http.ResponseWriter) {
	events, err := h.Repo.GetAllEvents()
	if err != nil {
		utils.HandleError("Failed to retrieve events from DB. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

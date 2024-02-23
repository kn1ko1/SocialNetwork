package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Endpoint: /api/notifications
// Allowed methods: POST

type NotificationsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewNotificationsHandler(r repo.IRepository) *NotificationsHandler {
	return &NotificationsHandler{Repo: r}
}

func (h *NotificationsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *NotificationsHandler) post(w http.ResponseWriter, r *http.Request) {

	var notification models.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received notification:", notification.NotificationType)

	// // Validate the event
	if validationErr := notification.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.CreateNotification(notification)
	if createErr != nil {
		utils.HandleError("Failed to create notification in the repository:", createErr)
		http.Error(w, "Failed to create notification", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Notification created successfully!"))
}

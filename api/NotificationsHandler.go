package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
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
	ctime := time.Now().UTC().UnixMilli()
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	notification.CreatedAt = ctime
	notification.UpdatedAt = ctime
	log.Println("[api/NotificationsHandler] Received notification:", notification)

	// // Validate the event
	if validationErr := notification.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	result, err := h.Repo.CreateNotification(notification)
	if err != nil {
		utils.HandleError("Failed to create notification in the repository:", err)
		http.Error(w, "Failed to create notification", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

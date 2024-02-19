package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET

type NotificationByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewNotificationByUserIdHandler(r repo.IRepository) *NotificationByUserIdHandler {
	return &NotificationByUserIdHandler{Repo: r}
}

func (h *NotificationByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *NotificationByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userId, userIdrr := strconv.Atoi(fields[len(fields)-1])
	if userIdrr != nil {
		utils.HandleError("Problem with AtoI userId. ", userIdrr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	notification, err := h.Repo.GetNotificationsByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get notification in GetNotificationByUserIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(notification)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"time"
)

// Endpoint: /api/GroupUser/
// Allowed methods: GET, PUT, DELETE

type GroupUsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUsersHandler(r repo.IRepository) *GroupUsersHandler {
	return &GroupUsersHandler{Repo: r}
}

func (h *GroupUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GroupUsersHandler) post(w http.ResponseWriter, r *http.Request) {

	var groupUser models.GroupUser
	err := json.NewDecoder(r.Body).Decode(&groupUser)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	ctime := time.Now().UTC().UnixMilli()
	groupUser.CreatedAt = ctime
	groupUser.UpdatedAt = ctime

	// Validate the groupUser
	if validationErr := groupUser.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create groupUser in the repository
	result, createErr := h.Repo.CreateGroupUser(groupUser)
	if createErr != nil {
		utils.HandleError("Failed to create groupUser in the repository:", createErr)
		http.Error(w, "Failed to create groupUser", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		// Note: Do not return an HTTP error here, as the resource has already been created successfully.
		// You may log the error or handle it as needed.
		return
	}

}

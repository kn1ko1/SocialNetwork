package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
)

type PrivacyHandler struct {
	Repo repo.IRepository
}

func NewPrivacyHandler(r repo.IRepository) *PrivacyHandler {
	return &PrivacyHandler{Repo: r}
}

func (h *PrivacyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodGet:
		h.UpdatePrivacyStatus(w, r)
		return
	case http.MethodPut:
		h.UpdatePrivacyStatus(w, r) // or any other method you want to call for PUT
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)

		return
	}
}

func (h *PrivacyHandler) UpdatePrivacyStatus(w http.ResponseWriter, r *http.Request) {
	var updateRequest struct {
		UserId   int  `json:"userId"`
		IsPublic bool `json:"isPublic"`
	}
	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = h.Repo.UpdateIsPublic(updateRequest.UserId, updateRequest.IsPublic)
	if err != nil {
		http.Error(w, "Failed to update privacy status", http.StatusInternalServerError)
		return
	}

	log.Printf("Updating privacy status for user %d to %t", updateRequest.UserId, updateRequest.IsPublic)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Privacy status updated successfully"})
}

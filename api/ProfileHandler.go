package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/profile  ?
// Allowed methods: GET

type ProfileHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewProfileHandler(r repo.IRepository) *ProfileHandler {
	return &ProfileHandler{Repo: r}
}

// A ProfileHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *ProfileHandler) get(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	userIdStr := fields[len(fields)-1]

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		user, userErr := getUser(r)
		if userErr != nil {
			utils.HandleError("Problem getting user.", userErr)
		}
		userId = user.UserId
	}

	profileData, err := h.Repo.GetProfileDataForUser(userId)
	if err != nil {
		utils.HandleError("Failed to get profileData in ProfileData. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(profileData)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *ProfileHandler) UpdatePrivacyStatus(w http.ResponseWriter, r *http.Request) {
// 	// Decode the request body to get the updated privacy status
// 	var updateRequest struct {
// 		UserId   int  `json:"userId"`
// 		IsPublic bool `json:"isPublic"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&updateRequest)
// 	if err != nil {
// 		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Call the repository method to update the IsPublic field
// 	err = h.Repo.UpdateUserPrivacyStatus(updateRequest.UserId, updateRequest.IsPublic)
// 	if err != nil {
// 		http.Error(w, "Failed to update privacy status", http.StatusInternalServerError)
// 		return
// 	}
// 	// Update the user's privacy status in the database
// 	// This is a simplified example, you would typically use a database query to update the user's privacy status
// 	log.Printf("Updating privacy status for user %s to %t", updateRequest.UserId, updateRequest.IsPublic)

// 	// Send a success response
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "Privacy status updated successfully"})
// }

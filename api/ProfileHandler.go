package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
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

	user, userErr := getUser(r)
	if userErr != nil {
		utils.HandleError("Problem getting user.", userErr)
	}

	profileData, err := h.Repo.GetProfileDataForUser(user.UserId)
	if err != nil {
		utils.HandleError("Failed to get profileData in ProfileData. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	profileData.ProfileUserData = transport.ProfileRegistrationInfo{
		UserId:    user.UserId,
		Bio:       user.Bio,
		DOB:       user.DOB,
		Email:     user.Email,
		FirstName: user.FirstName,
		ImageURL:  user.ImageURL,
		IsPublic:  user.IsPublic,
		LastName:  user.LastName,
		Username:  user.Username,
	}

	err = json.NewEncoder(w).Encode(profileData)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

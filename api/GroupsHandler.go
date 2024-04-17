package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"
)

// Endpoint: /api/groups
// Allowed methods: GET, POST, PUT, DELETE

type GroupsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupsHandler(r repo.IRepository) *GroupsHandler {
	return &GroupsHandler{Repo: r}
}

// Supported Methods: GET, POST
func (h *GroupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *GroupsHandler) post(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("User unauthorized", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// contentType := r.Header.Get("Content-Type")
	var group models.Group
	// switch contentType {
	// case "application/json":
	// err := json.NewDecoder(r.Body).Decode(&group)
	// if err != nil {
	// 	utils.HandleError("Failed to decode request body:", err)
	// 	http.Error(w, "Failed to decode request body", http.StatusBadRequest)
	// 	return
	// }
	// case "application/x-www-form-urlencoded":
	// 	err := r.ParseForm()
	// 	if err != nil {
	// 		utils.HandleError("Failed to parse form:", err)
	// 		http.Error(w, "internal server error", http.StatusInternalServerError)
	// 		return
	// 	}
	ctime := time.Now().UTC().UnixMilli()
	group.CreatedAt = ctime
	log.Println("This createdat:", group.CreatedAt)
	group.CreatorId = user.UserId
	group.Description = r.PostFormValue("group-description")
	log.Println("this group description:", group.Description)
	group.Title = r.PostFormValue("group-title")
	log.Println("this group title:", group.Title)
	group.UpdatedAt = ctime
	//}
	// Validate the group
	if validationErr := group.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}
	log.Println("Received group:", group.Title, group.Description)
	// Create group in the repository
	result, createErr := h.Repo.CreateGroup(group)
	if createErr != nil {
		utils.HandleError("Failed to create group in the repository:", createErr)
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
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

func (h *GroupsHandler) get(w http.ResponseWriter) {
	allGroups, err := h.Repo.GetAllGroups()
	if err != nil {
		utils.HandleError("Failed to get group in GroupHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&allGroups)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *GroupsHandler) put(w http.ResponseWriter, r *http.Request) {

// 	var group models.Group
// 	err := json.NewDecoder(r.Body).Decode(&group)
// 	if err != nil {
// 		utils.HandleError("Failed to decode request body:", err)
// 		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("Received group:", group.Title, group.Description)

// 	// Validate the group
// 	if validationErr := group.Validate(); validationErr != nil {
// 		utils.HandleError("Validation failed:", validationErr)
// 		http.Error(w, "Validation failed", http.StatusBadRequest)
// 		return
// 	}

// 	// Create group in the repository
// 	result, createErr := h.Repo.UpdateGroup(group)
// 	if createErr != nil {
// 		utils.HandleError("Failed to update group in the repository:", createErr)
// 		http.Error(w, "Failed to update group", http.StatusInternalServerError)
// 		return
// 	}

// 	// Encode and write the response
// 	err = json.NewEncoder(w).Encode(result)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	// Correct HTTP header for a newly created resource:
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("Post updated successfully!"))
// }

// func (h *GroupsHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	err := h.Repo.DeleteAllGroups()
// 	if err != nil {
// 		utils.HandleError("Failed to delete all groups. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("groups were deleted"))
// }

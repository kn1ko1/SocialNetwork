package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, PUT, DELETE

type GroupByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupByIdHandler(r repo.IRepository) *GroupByIdHandler {
	return &GroupByIdHandler{Repo: r}
}

// A GroupsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GroupByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	groupId, GroupIdErr := strconv.Atoi(fields[len(fields)-1])
	if GroupIdErr != nil {
		utils.HandleError("Problem with AtoI GroupId. ", GroupIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Group, err := h.Repo.GetGroupById(groupId)
	if err != nil {
		utils.HandleError("Failed to get Group in GetGroupByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(Group)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GroupByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var Group models.Group
	err := json.NewDecoder(r.Body).Decode(&Group)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Group to update:", Group.Title)

	// Validate the Group
	if validationErr := Group.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateGroup(Group)
	if createErr != nil {
		utils.HandleError("Failed to update group in the repository:", createErr)
		http.Error(w, "Failed to update group", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GroupByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	eventIdStr := fields[len(fields)-1]
	groupId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for GroupId:", groupId)

	err = h.Repo.DeleteGroup(groupId)
	if err != nil {
		utils.HandleError("Failed to delete Groups. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

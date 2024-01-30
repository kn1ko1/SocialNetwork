package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
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
	case http.MethodDelete:
		h.delete(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GroupByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	GroupIdString := queryParams.Get("GroupId")
	GroupId, GroupIdErr := strconv.Atoi(GroupIdString)
	if GroupIdErr != nil {
		log.Println("Problem with AtoI GroupId. ", GroupIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	Group, err := h.Repo.GetGroup(GroupId)
	if err != nil {
		log.Println("Failed to get Group in GetGroupByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(Group)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GroupByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var Group models.Group
	err := json.NewDecoder(r.Body).Decode(&Group)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Group to update:", Group.Title)

	// Validate the Group
	if validationErr := Group.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateGroup(Group)
	if createErr != nil {
		log.Println("Failed to update group in the repository:", createErr)
		http.Error(w, "Failed to update group", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *GroupByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	GroupIdString := queryParams.Get("GroupId")
	GroupId, GroupIdErr := strconv.Atoi(GroupIdString)
	if GroupIdErr != nil {
		log.Println("Problem with AtoI GroupId. ", GroupIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for GroupId:", GroupId)

	err := h.Repo.DeleteGroup(GroupId)
	if err != nil {
		log.Println("Failed to delete Groups. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

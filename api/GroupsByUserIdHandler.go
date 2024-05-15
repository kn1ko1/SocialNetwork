package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, PUT, DELETE

type GroupsByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupsByUserIdHandler(r repo.IRepository) *GroupsByUserIdHandler {
	return &GroupsByUserIdHandler{Repo: r}
}

// A GroupsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupsByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *GroupsByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userId, UserIdErr := strconv.Atoi(fields[len(fields)-2])
	if UserIdErr != nil {
		utils.HandleError("Problem with AtoI UserId. ", UserIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	groupUsers, err := h.Repo.GetGroupUsersByUserId(userId)
	if err != nil {
		utils.HandleError("Failed to get Group in GetGroupsByUserIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var groups []models.Group
	for i := 0; i < len(groupUsers); i++ {
		group, err := h.Repo.GetGroupById(groupUsers[i].GroupId)
		if err != nil {
			utils.HandleError("Failed to get Group in GetGroupById. ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		}
		groups = append(groups, group)
	}

	err = json.NewEncoder(w).Encode(groups)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/groups/{groupId}/posts
// Allowed methods: GET, DELETE

type PostsByGroupIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsByGroupIdHandler(r repo.IRepository) *PostsByGroupIdHandler {
	return &PostsByGroupIdHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsByGroupIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByGroupIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	groupId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupPosts, err := h.Repo.GetPostsByGroupId(groupId)
	if err != nil {
		utils.HandleError("Failed to get posts in PostsByGroupIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(groupPosts)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *PostsByGroupIdHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	// figure out postId
// 	queryParams := r.URL.Query()
// 	groupIdString := queryParams.Get("groupId")
// 	groupId, postIdErr := strconv.Atoi(groupIdString)
// 	if postIdErr != nil {
// 		utils.HandleError("Problem with AtoI groupId. ", postIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	utils.HandleError("Received delete request for groupId:", groupId)

// 	err := h.Repo.DeletePostByGroupId(groupId)
// 	if err != nil {
// 		utils.HandleError("Failed to delete Post. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

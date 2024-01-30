package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"strconv"
	"strings"
)

// Endpoint: /api/groups/{groupId}/posts
// Allowed methods: GET, DELETE

type GroupPostsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupPostsHandler(r repo.IRepository) *GroupPostsHandler {
	return &GroupPostsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *GroupPostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *GroupPostsHandler) get(w http.ResponseWriter, r *http.Request) {
	_, err := getUser(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	fields := strings.Split(r.URL.Path, "/")
	groupId, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupPosts, err := h.Repo.GetPostsByGroupId(groupId)
	if err != nil {
		log.Println("Failed to get posts in GroupPostsHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(groupPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *GroupPostsHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	// figure out postId
// 	queryParams := r.URL.Query()
// 	groupIdString := queryParams.Get("groupId")
// 	groupId, postIdErr := strconv.Atoi(groupIdString)
// 	if postIdErr != nil {
// 		log.Println("Problem with AtoI groupId. ", postIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Received delete request for groupId:", groupId)

// 	err := h.Repo.DeletePostByGroupId(groupId)
// 	if err != nil {
// 		log.Println("Failed to delete Post. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: get/api/groupId/
// Allowed methods: GET

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
	// Get Session Cookie
	c, err := r.Cookie("Session")
	if err != nil {
		// Log Error
		log.Println(err.Error())
		// Return HTTP Status Unauthorized
		//
		// N.B. for simplicity of the example, we are simply returning
		// an HTTP error. In the actual project, probably a JSON payload.
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Authenticate Session Cookie - user variable discarded because user struct not used here...
	_, err = auth.AuthenticateSessionCookie(c)
	if err != nil {
		// Same error as above - maker of request is unauthorized
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByGroupIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	groupIdString := queryParams.Get("groupId")
	groupId, groupIdErr := strconv.Atoi(groupIdString)
	if groupIdErr != nil {
		log.Println("Problem with AtoI senderId. ", groupIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	groupPosts, err := h.Repo.GetPostsByGroupId(groupId)
	if err != nil {
		log.Println("Failed to get posts in PostsByGroupIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(groupPosts)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
}

func (h *PostsByGroupIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out postId
	queryParams := r.URL.Query()
	groupIdString := queryParams.Get("groupId")
	groupId, postIdErr := strconv.Atoi(groupIdString)
	if postIdErr != nil {
		log.Println("Problem with AtoI groupId. ", postIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for groupId:", groupId)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeletePostByGroupId(groupId)
	if err != nil {
		log.Println("Failed to delete Post. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("post was deleted"))
}

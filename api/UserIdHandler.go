package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
)

// Allowed methods: POST

type UserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserIdHandler(r repo.IRepository) *UserIdHandler {
	return &UserIdHandler{Repo: r}
}

// A UserUsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserIdHandler) get(w http.ResponseWriter, r *http.Request) {

	user, err := getUser(r)
	if err != nil {
		utils.HandleError("Failed to get user for UserId. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("I made it to UserIdHandler")
	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user.UserId)
	log.Println("api/UserIdHandler Current Is is: ", user.UserId)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

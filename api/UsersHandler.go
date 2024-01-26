package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/users
// Allowed methods: GET, POST

type UsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUsersHandler(r repo.IRepository) *UsersHandler {
	return &UsersHandler{Repo: r}
}

// A UsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UsersHandler) post(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received user:", user.Username)

	// Validate the user
	if validationErr := user.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.CreateUser(user)
	if createErr != nil {
		log.Println("Failed to create user in the repository:", createErr)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *UsersHandler) get(w http.ResponseWriter, r *http.Request) {

	userUsers, err := h.Repo.GetAllUsers()
	if err != nil {
		log.Println("Failed to get all Users. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userUsers)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

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
	// case http.MethodGet:
	// 	h.get(w, r)
	// 	return
	// case http.MethodPut:
	// 	h.put(w, r)
	// 	return
	// case http.MethodDelete:
	// 	h.delete(w, r)
	// 	return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Separate out HTTP methods for clean separation of concerns
// N.B. Use lowercase names, i.e. "post", "get", etc. for correct encapsulation
func (h *UsersHandler) post(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received user:", user.Username)

	// user := models.User{
	// 	CreatedAt:         111111111,
	// 	DOB:               2221111,
	// 	Email:             "example@example.com",
	// 	EncryptedPassword: "eXaMpLe",
	// 	FirstName:         "Rupert",
	// 	IsPublic:          true,
	// 	LastName:          "Cheetham",
	// 	UpdatedAt:         111111111,
	// 	Username:          "Ardek"}

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
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully!"))
}

// func (h *UsersHandler) get(w http.ResponseWriter, r *http.Request) {

// 	allUsers, err := h.Repo.GetAllUsers()
// 	if err != nil {
// 		log.Println("Failed to get all users in UserHandler. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(allUsers)
// 	if err != nil {
// 		log.Println("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Here are all users"))
// }

// func (h *UsersHandler) put(w http.ResponseWriter, r *http.Request) {

// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		log.Println("Failed to decode request body:", err)
// 		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("Updating User:", user.UserId, user.Username)

// 	// Example User to test function
// 	// user := models.User{
// 	// Bio: "Bio (update)"
// 	// 	CreatedAt:         111111111,
// 	// 	DOB:               2221111,
// 	// 	Email:             "example@example.com",
// 	// 	EncryptedPassword: "eXaMpLe",
// 	// 	FirstName:         "Rupert",
// 	// 	IsPublic:          true,
// 	// 	LastName:          "Cheetham",
// 	// 	UpdatedAt:         3333333333,
// 	// 	Username:          "Ardek"}

// 	// Validate the post
// 	if validationErr := user.Validate(); validationErr != nil {
// 		log.Println("Validation failed:", validationErr)
// 		http.Error(w, "Validation failed", http.StatusBadRequest)
// 		return
// 	}

// 	// Create post in the repository
// 	result, createErr := h.Repo.UpdateUser(user)
// 	if createErr != nil {
// 		log.Println("Failed to update user in the repository:", createErr)
// 		http.Error(w, "Failed to update user", http.StatusInternalServerError)
// 		return
// 	}

// 	// Encode and write the response
// 	err = json.NewEncoder(w).Encode(result)
// 	if err != nil {
// 		log.Println("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	// Correct HTTP header for a newly created resource:
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("User updated successfully!"))
// }

// func (h *UsersHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	err := h.Repo.DeleteAllUsers()
// 	if err != nil {
// 		log.Println("Failed to delete all User. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("user deleted successfully"))
// }

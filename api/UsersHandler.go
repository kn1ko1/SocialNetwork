package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/users
// Allowed methods: GET, POST, PUT, Delete

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

	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
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

func (h *UsersHandler) get(w http.ResponseWriter, r *http.Request) {
	// Not Implemented - would be h.Repo.GetAllUsers() ... you get the idea
	w.Write([]byte("Here are your users!"))
}

func (h *UsersHandler) put(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating User:", user.UserId, user.Username)

	// Example User to test function
	// user := models.User{
	// Bio: "Bio (update)"
	// 	CreatedAt:         111111111,
	// 	DOB:               2221111,
	// 	Email:             "example@example.com",
	// 	EncryptedPassword: "eXaMpLe",
	// 	FirstName:         "Rupert",
	// 	IsPublic:          true,
	// 	LastName:          "Cheetham",
	// 	UpdatedAt:         3333333333,
	// 	Username:          "Ardek"}

	// Validate the post
	if validationErr := user.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.UpdateUser(user)
	if createErr != nil {
		log.Println("Failed to update user in the repository:", createErr)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
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
	w.Write([]byte("User updated successfully!"))
}

func (h *UsersHandler) delete(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// figure out postId
	var userId int
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received delete request for userId:", userId)

	// example postId for testing
	// postId := 1

	err = h.Repo.DeleteUserById(userId)
	if err != nil {
		log.Println("Failed to delete User. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted successfully"))
}

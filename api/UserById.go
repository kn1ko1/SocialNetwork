package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: /api/users/user/{userId}
// Allowed methods: GET, DELETE

type UserByUserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserByUserIdHandler(r repo.IRepository) *UserByUserIdHandler {
	return &UserByUserIdHandler{Repo: r}
}

// A UsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserByUserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserByUserIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	postIdString := queryParams.Get("userId")
	userId, userIdErr := strconv.Atoi(postIdString)
	if userIdErr != nil {
		log.Println("Problem with AtoI userId. ", userId)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userUsers, err := h.Repo.GetUserById(userId)
	if err != nil {
		log.Println("Failed to get Users in GetUserByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userUsers)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your Users"))
}

func (h *UserByUserIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating:", user.UserId, user.Username)

	// Example User to test function
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

	// Validate the User <3
	if validationErr := user.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update user in the repository
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
	w.Write([]byte("Post updated successfully!"))
}

func (h *UserByUserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		log.Println("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeleteUserById(userID)
	if err != nil {
		log.Println("Failed to delete Users. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Users were deleted"))
}

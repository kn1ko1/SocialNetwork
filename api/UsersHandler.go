package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"

	"golang.org/x/crypto/bcrypt"
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
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received user:", user.Username)

	// Validate the user
	if validationErr := user.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleError("error in encryption", err)
	}

	user.EncryptedPassword = string(hashPassword)

	parseMultipartFormErr := r.ParseMultipartForm(10 << 20)
	if parseMultipartFormErr != nil {
		utils.HandleError("Unable to Parse Multipart Form.", parseMultipartFormErr)
	}

	file, fileHeader, formFileErr := r.FormFile("image")
	if formFileErr != nil {
		utils.HandleError("Error reading image.", formFileErr)
	}

	defer file.Close()

	//if file is given
	if file != nil {
		var imageHandlerErr error
		user.ImageURL, imageHandlerErr = ImageProcessing(w, r, file, *fileHeader)
		if imageHandlerErr != nil {
			utils.HandleError("Error with ImageHandler", imageHandlerErr)
		}
		fmt.Println("USER INSERTED WITH FILE")
	} else {
		fmt.Println("USER INSERTED WITHOUT FILE")
	}

	// Create post in the repository
	result, createErr := h.Repo.CreateUser(user)
	if createErr != nil {
		utils.HandleError("Failed to create user in the repository:", createErr)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *UsersHandler) get(w http.ResponseWriter, r *http.Request) {

	userUsers, err := h.Repo.GetAllUsers()
	if err != nil {
		utils.HandleError("Failed to get all Users. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

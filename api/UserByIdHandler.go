package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/users/{userId}
// Allowed methods: GET, PUT, DELETE

type UserByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewUserByIdHandler(r repo.IRepository) *UserByIdHandler {
	return &UserByIdHandler{Repo: r}
}

// A UsersHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *UserByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *UserByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	userId, userIdErr := strconv.Atoi(fields[len(fields)-1])
	if userIdErr != nil {
		utils.HandleError("Problem with AtoI userId. ", userIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userUsers, err := h.Repo.GetUserById(userId)
	if err != nil {
		utils.HandleError("Failed to get Users in GetUserById. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userUsers)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here is the User"))
}

func (h *UserByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Updating:", user.UserId, user.Username)

	// Validate the User <3
	if validationErr := user.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

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
	}
	// Update user in the repository
	result, createErr := h.Repo.UpdateUser(user)
	if createErr != nil {
		utils.HandleError("Failed to update user in the repository:", createErr)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User updated successfully!"))
}

func (h *UserByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		utils.HandleError("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	err := h.Repo.DeleteUserById(userID)
	if err != nil {
		utils.HandleError("Failed to delete Users. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

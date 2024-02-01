package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationHandler struct {
	Repo repo.IRepository
}

func NewRegistrationHandler(r repo.IRepository) *RegistrationHandler {
	return &RegistrationHandler{Repo: r}
}

func (h *RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *RegistrationHandler) post(w http.ResponseWriter, r *http.Request) {
	// Enable CORS headers for this handler
	// handlers.SetupCORS(&w, r)

	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.HandleError("Error with Registration.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.Validate()

	if err != nil {
		http.Error(w, "validation failed for user registration", http.StatusBadRequest)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleError("Error with password encryption", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.EncryptedPassword = string(hashPassword)

	log.Println("Received user:", user)
	// _, err = sqlite.CreateUser(db, user)
	if err != nil {
		utils.HandleError("Unable to register a new user in AddUserHandler", err)
		http.Error(w, "Unable to register a new user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

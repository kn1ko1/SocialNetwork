package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
	"time"

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
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *RegistrationHandler) post(w http.ResponseWriter, r *http.Request) {
	var user transport.RegisteringUser
	ctime := time.Now().UTC().UnixMilli()

	var processedUser models.User

	cookie, err := r.Cookie(cookieName)
	if err == nil {
		_, exists := sessionMap[cookie.Value]
		if exists {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}

	json.NewDecoder(r.Body).Decode(&user)

	log.Println("[RegistrationHandler] ctime:", ctime)
	user.CreatedAt = ctime

	// Parse the date string into a time.Time object
	date, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Println("date", date)
	fmt.Println(user.DOB)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleError("Error with password encryption", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	user.EncryptedPassword = string(encryptedPassword)
	user.UpdatedAt = ctime

	processedUser.Bio = user.Bio
	processedUser.CreatedAt = ctime
	processedUser.DOB = date.UnixNano() / int64(time.Millisecond)
	processedUser.Email = user.Email
	processedUser.EncryptedPassword = string(encryptedPassword)
	processedUser.FirstName = user.FirstName
	processedUser.ImageURL = user.ImageURL
	processedUser.IsPublic = user.IsPublic
	processedUser.LastName = user.LastName
	processedUser.UpdatedAt = ctime
	processedUser.Username = user.Username
	//err = user.Validate()

	if err != nil {
		utils.HandleError("User invalid", err)
		http.Error(w, "validation failed for user registration", http.StatusBadRequest)
		return
	}
	log.Println("Received user in RegistrationHandler:", user)

	user, err = h.Repo.CreateUser(processedUser)
	if err != nil {
		utils.HandleError("Unable to register a new user in AddUserHandler", err)
		http.Error(w, "Unable to register a new user", http.StatusBadRequest)
		return
	}
	cookieValue = GenerateNewUUID()
	//sessionMap[cookieValue] = &user

	cookie = &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Path:     "/",
		Expires:  time.Now().Add(timeout),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

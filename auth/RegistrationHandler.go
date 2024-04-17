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

	// Checks cookies
	// cookie, err := r.Cookie(CookieName)
	// if err == nil {
	// 	_, exists := SessionMap[cookie.Value]
	// 	if exists {
	// 		utils.HandleError("Login failed - user already logged in:", err)
	// 		http.Error(w, "user already logged in", http.StatusBadRequest)
	// 		return
	// 	}
	// }
	cookie, err := r.Cookie("SessionID")
	if err == nil {
		_, err = DefaultManager.Get(cookie.Value)
		if err == nil {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}

	// Decodes incoming json into registeringUser (transport.RegisteringUser)
	var registeringUser transport.RegisteringUser
	json.NewDecoder(r.Body).Decode(&registeringUser)

	// converts date from string to milliseconds for storage in database
	date, err := time.Parse("2006-01-02", registeringUser.DOB)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	dateInMilliseconds := date.UTC().UnixMilli()

	// Encrypt Password for Storage
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(registeringUser.EncryptedPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleError("Error with password encryption", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// Creates current time in milliseconds for CreatedAt and UpdatedAt fields
	ctime := time.Now().UTC().UnixMilli()

	// Creates a models.User struct to passing to CreateUser, so that it both takes in and passes out the same types of data
	processedUser := models.User{
		Bio:               registeringUser.Bio,
		CreatedAt:         ctime,
		DOB:               dateInMilliseconds,
		Email:             registeringUser.Email,
		EncryptedPassword: string(encryptedPassword),
		FirstName:         registeringUser.FirstName,
		ImageURL:          registeringUser.ImageURL,
		IsPublic:          registeringUser.IsPublic,
		LastName:          registeringUser.LastName,
		UpdatedAt:         ctime,
		Username:          registeringUser.Username,
	}

	err = processedUser.Validate()
	if err != nil {
		utils.HandleError("User invalid", err)
		http.Error(w, "validation failed for user registration", http.StatusBadRequest)
		return
	}
	log.Println("Received user in RegistrationHandler:", processedUser)

	processedUser, err = h.Repo.CreateUser(processedUser)
	if err != nil {
		utils.HandleError("Unable to register a new user in AddUserHandler", err)
		http.Error(w, "Unable to register a new user", http.StatusBadRequest)
		return
	}

	// Construct the response JSON
	response := struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		User    interface{} `json:"user,omitempty"`
	}{
		Success: true,
		Message: "Registration successful",
		User: struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			// Add other user fields as needed
		}{
			ID:       processedUser.UserId,
			Username: processedUser.Username,
			Email:    processedUser.Email,
			// Assign other user fields as needed
		},
	}

	// Sets up a new Cookie
	cookieValue, err := generateCookieValue()
	if err != nil {
		utils.HandleError("Cookie Generation Error", err)
		http.Error(w, "Problem generating cookie", http.StatusInternalServerError)
		return
	}
	err = DefaultManager.Add(cookieValue, processedUser)
	if err != nil {
		utils.HandleError("Cookie Generation Error", err)
		http.Error(w, "Problem generating cookie", http.StatusInternalServerError)
		return
	}
	c := http.Cookie{
		Name:     "SessionID",
		Value:    cookieValue,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   DefaultManager.Lifetime(),
	}
	// fmt.Println(c.Value)
	http.SetCookie(w, &c)
	// Convert the response struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.HandleError("Failed to marshal JSON response", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		utils.HandleError("Failed to write JSON response", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

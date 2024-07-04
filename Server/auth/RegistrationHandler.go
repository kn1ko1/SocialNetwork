package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"socialnetwork/Server/imageProcessing"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
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

	cookie, err := r.Cookie("SessionID")
	if err == nil {
		_, err = DefaultManager.Get(cookie.Value)
		if err == nil {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}

	ctime := time.Now().UTC().UnixMilli()
	// Parse form data
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.HandleError("Failed to parse form data:", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract form fields
	bio := r.FormValue("bio")
	date, err := time.Parse("2006-01-02", r.FormValue("dob"))
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	dateInMilliseconds := date.UTC().UnixMilli()
	email := r.FormValue("email")
	password := r.FormValue("password")
	log.Println("password", password)
	// Encrypt Password for Storage
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.HandleError("Error with password encryption", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	firstName := r.FormValue("firstName")
	imageURL := ""
	file, fileHeader, _ := r.FormFile("image")
	if file != nil {

		defer file.Close()
		imageURL, err = imageProcessing.ImageProcessing(w, r, file, *fileHeader)
		if err != nil {
			utils.HandleError("Error with ImageHandler", err)
			// http.Error(w, "Failed to process image", http.StatusInternalServerError)
			return
		}
		log.Println("[api/PostsHandler] Image Stored at:", imageURL)
	}
	isPublic := r.FormValue("isPublic")
	isPublicBool := false
	if isPublic == "true" {
		isPublicBool = true
	}
	lastName := r.FormValue("lastName")
	username := r.FormValue("username")

	user := models.User{
		Bio:               bio,
		CreatedAt:         ctime,
		DOB:               dateInMilliseconds,
		Email:             email,
		EncryptedPassword: string(encryptedPassword),
		FirstName:         firstName,
		ImageURL:          imageURL,
		IsPublic:          isPublicBool,
		LastName:          lastName,
		UpdatedAt:         ctime,
		Username:          username,
	}

	err = user.Validate()
	if err != nil {
		utils.HandleError("User invalid", err)
		http.Error(w, "validation failed for user registration", http.StatusBadRequest)
		return
	}
	log.Println("Received user in RegistrationHandler:", user)

	user, err = h.Repo.CreateUser(user)
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
			ID:       user.UserId,
			Username: user.Username,
			Email:    user.Email,
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
	err = DefaultManager.Add(cookieValue, user)
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

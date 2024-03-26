package auth

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Repo repo.IRepository
}

func NewLoginHandler(r repo.IRepository) *LoginHandler {
	return &LoginHandler{Repo: r}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *LoginHandler) post(w http.ResponseWriter, r *http.Request) {
	var loginInfo transport.LoginInfo

	json.NewDecoder(r.Body).Decode(&loginInfo)

	user, err := h.Repo.GetUserByUsernameOrEmail(loginInfo.UsernameOrEmail)
	if err != nil {
		utils.HandleError("Failed to retrieve user", err)
		http.Error(w, "user with specified username or email does not exist", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(loginInfo.Password))
	if err != nil {
		utils.HandleError("Failed to retrieve user", err)
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	// If login is successful, construct the response JSON
	response := struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		User    interface{} `json:"user,omitempty"`
	}{
		Success: true,
		Message: "Login successful",
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

	SessionMap[CookieValue] = &user

	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    CookieValue,
		Path:     "/",
		Expires:  time.Now().Add(timeout),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
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

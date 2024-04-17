package auth

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	rp repo.IRepository
}

func NewLoginHandler(r repo.IRepository) *LoginHandler {
	ret := new(LoginHandler)
	ret.rp = r
	return ret
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *LoginHandler) post(w http.ResponseWriter, r *http.Request) {
	var loginInfo transport.LoginInfo
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	user, err := h.rp.GetUserByUsernameOrEmail(loginInfo.UsernameOrEmail)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(loginInfo.Password))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}
	cookieValue, err := generateCookieValue()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusBadRequest)
		return
	}
	err = DefaultManager.Add(cookieValue, user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
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

func generateCookieValue() (string, error) {
	var val [64]byte
	for i := 0; i < 64; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return "", err
		}
		val[i] = byte(random.Int64())
	}
	return base64.URLEncoding.EncodeToString(val[:]), nil
}

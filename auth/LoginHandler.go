package auth

import (
	"encoding/json"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
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
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *LoginHandler) post(w http.ResponseWriter, r *http.Request) {

	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// usr, err := sqlite.GetUserById(db, user.UserId)

	if err != nil {
		http.Error(w, "Unable to get UserId", http.StatusBadRequest)
		return
	}

	// Set the session ID as a cookie
	sessionCookie := http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Expires:  SessionExpiration,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   int(timeout.Seconds()),
		SameSite: http.SameSiteNoneMode,
	}

	sessionMap[sessionCookie.Value] = user
	reflectedSessionMap[user] = sessionCookie.Value

	http.SetCookie(w, &sessionCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// json.NewEncoder(w).Encode(usr)
}

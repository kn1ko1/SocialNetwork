package auth

import (
	"encoding/json"
	"net/http"
	"socialnetwork/repo"
	"strings"
	"time"
)

// Logout Handler

type LogoutHandler struct {
	Repo repo.IRepository
}

func NewLogoutHandler(r repo.IRepository) *LogoutHandler {
	return &LogoutHandler{Repo: r}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *LogoutHandler) post(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:     CookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   strings.HasPrefix(r.Header.Get("X-Forwarded-Proto"), "https"),
		Domain:   "localhost",
		SameSite: http.SameSiteNoneMode, // Set SameSite attribute to None
	}

	http.SetCookie(w, &cookie)

	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{
		Success: true,
		Message: "Logout successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send the response
	json.NewEncoder(w).Encode(response)
}

package auth

import (
	"net/http"
	"socialnetwork/repo"
	"time"
)

// Logout Handler

type LogoutHandler struct {
	Repo repo.IRepository
}

func NewLogoutHandler(r repo.IRepository) *LoginHandler {
	return &LoginHandler{Repo: r}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *LogoutHandler) get(w http.ResponseWriter, r *http.Request) {
	// handlers.SetupCORS(&w, r)
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		Domain:   "localhost",
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

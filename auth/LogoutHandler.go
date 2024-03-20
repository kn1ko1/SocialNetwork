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
	case http.MethodPost:
		h.post(w)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *LogoutHandler) post(w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:     CookieName,
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

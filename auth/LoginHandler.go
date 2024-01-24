package auth

import (
	"net/http"
	"socialnetwork/repo"
)

type LoginHandler struct {
	Repo repo.IRepository
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

	// Set the session ID as a cookie
	sessionCookie := http.Cookie{
		Name:     cookieName,
		Value:    "Session",
		Expires:  SessionExpiration,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   int(timeout.Seconds()),
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &sessionCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// jsonResponse, err := json.Marshal(msg)

	// if err != nil {
	// 	http.Error(w, "Internal sever error", http.StatusBadRequest)
	// 	return
	// }

	// w.Write(jsonResponse)
	// json.NewEncoder(w).Encode(user)
}

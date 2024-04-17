package auth

import (
	"log"
	"net/http"
	"socialnetwork/repo"
)

type LogoutHandler struct {
	rp repo.IRepository
}

func NewLogoutHandler(r repo.IRepository) *LogoutHandler {
	ret := new(LogoutHandler)
	ret.rp = r
	return ret
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *LogoutHandler) get(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	DefaultManager.Delete(c.Value)
	c = &http.Cookie{
		Name:     "Session",
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

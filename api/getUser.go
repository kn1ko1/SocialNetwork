package api

import (
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/utils"
)

// Obtains user by reading the cookie
func getUser(r *http.Request) (*models.User, error) {

	cookie, err := r.Cookie(auth.CookieName)
	if err != nil {
		utils.HandleError("Error reading cookie.", err)
		return nil, err
	}
	log.Println("[api/getUser] Cookie:", cookie)

	ret, err := auth.AuthenticateSessionCookie(cookie)
	if err != nil {
		utils.HandleError("Error authenticating session cookie.", err)
		return nil, err
	}

	return ret, nil
}

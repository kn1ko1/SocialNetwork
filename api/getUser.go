package api

import (
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/utils"
)

func getUser(r *http.Request) (*models.User, error) {
	c, err := r.Cookie("Session")
	if err != nil {
		utils.HandleError("Error reading cookie", err)
		return nil, err
	}
	ret, err := auth.AuthenticateSessionCookie(c)
	if err != nil {
		utils.HandleError("Error authenticating session cookie.", err)
		return nil, err
	}
	return ret, nil
}

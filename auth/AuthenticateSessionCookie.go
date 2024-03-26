package auth

import (
	"errors"
	"log"
	"net/http"
	"socialnetwork/models"
)

func AuthenticateSessionCookie(c *http.Cookie) (*models.User, error) {

	user, exists := SessionMap[c.Value]
	log.Println("[AuthenticateSessionCookie] User is: ", user)
	if !exists {
		return nil, errors.New("session cookie invalid")
	}
	return user, nil
}

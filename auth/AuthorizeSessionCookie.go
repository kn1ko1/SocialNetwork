package auth

import (
	"errors"
	"net/http"
	"socialnetwork/models"
)

func AuthenticateSessionCookie(c *http.Cookie) (*models.User, error) {
	user, exists := sessionMap[c.Value]
	if !exists {
		return nil, errors.New("session cookie invalid")
	}
	return user, nil
}

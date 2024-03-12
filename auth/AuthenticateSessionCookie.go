package auth

import (
	"errors"
	"log"
	"net/http"
	"socialnetwork/models"
)

func AuthenticateSessionCookie(c *http.Cookie) (*models.User, error) {
	sessionMap := make(map[string]*models.User)
	testUser := models.GenerateValidUser()
	// testUser.UserId = rand.Intn(101)
	testUser.UserId = 1

	sessionMap[c.Value] = testUser

	user, exists := sessionMap[c.Value]
	log.Println("[AuthenticateSessionCookie] User is: ", user)
	log.Println("[AuthenticateSessionCookie] Currently fakes a user since login is not running")
	if !exists {
		return nil, errors.New("session cookie invalid")
	}
	return user, nil
}

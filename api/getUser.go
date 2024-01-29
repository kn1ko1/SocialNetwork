package api

import (
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
)

func getUser(r *http.Request) (*models.User, error) {
	c, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	ret, err := auth.AuthenticateSessionCookie(c)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return ret, nil
}

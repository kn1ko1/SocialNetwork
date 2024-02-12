package api

import (
	"math/rand"
	"net/http"
	"socialnetwork/models"
)

func getUser(r *http.Request) (*models.User, error) {
	// c, err := r.Cookie("Session")
	// if err != nil {
	// 	utils.HandleError("Error reading cookie", err)
	// 	return nil, err
	// }
	// ret, err := auth.AuthenticateSessionCookie(c)
	// if err != nil {
	// 	utils.HandleError("Error authenticating session cookie.", err)
	// 	return nil, err
	// }

	//
	// For testing only
	ret := models.GenerateValidUser()
	ret.UserId = rand.Intn(101)
	//
	//

	return ret, nil
}

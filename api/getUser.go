package api

import (
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/utils"
)

func getUser(r *http.Request) (*models.User, error) {

	// c, err := r.Cookie("Session")
	// if err != nil {
	// 	utils.HandleError("Error reading cookie", err)
	// 	return nil, err
	// }
	cookie := http.Cookie{
		Name:  "sessionID",
		Value: "fakeSessionID123",
	}

	ret, err := auth.AuthenticateSessionCookie(&cookie)
	if err != nil {
		utils.HandleError("Error authenticating session cookie.", err)
		return nil, err
	}

	//
	// For testing only
	// ret := models.GenerateValidUser()
	// ret.UserId = rand.Intn(101)
	//
	//
	log.Println("api/getUser.go.  UserId is:", ret.UserId)
	return ret, nil
}

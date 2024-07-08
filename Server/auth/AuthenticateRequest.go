package auth

import (
	"net/http"
	"socialnetwork/Server/models"
)

func AuthenticateRequest(r *http.Request) (models.User, error) {
	var ret models.User
	c, err := r.Cookie("SessionID")
	if err != nil {
		// log.Println(err.Error())
		return ret, err
	}
	ret, err = DefaultManager.Get(c.Value)
	if err != nil {
		// log.Println(err.Error())
		return ret, err
	}
	return ret, nil
}

package auth

import (
	"errors"
	"net/http"
	"socialnetwork/models"
)

var (
	SessionMap map[string]*models.User
)

func init() {
	SessionMap = make(map[string]*models.User)
}

// Not implemented - here we would authenticate the session cookie
// i.e. - look up in a map for a 'User' associated with this cookie value
// Can be done in a variety of ways
//
// For the example, I am just ommitting all this logic and simulating
// it successfully returning a 'Test' user
func AuthenticateSessionCookie(c *http.Cookie) (*models.User, error) {
	ret, exists := SessionMap[c.Value]
	if !exists {
		return nil, errors.New("session cookie invalid")
	}
	return ret, nil
}

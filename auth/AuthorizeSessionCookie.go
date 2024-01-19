package auth

import (
	"net/http"
	"socialnetwork/models"
)

// Not implemented - here we would authenticate the session cookie
// i.e. - look up in a map for a 'User' associated with this cookie value
// Can be done in a variety of ways
//
// For the example, I am just ommitting all this logic and simulating
// it successfully returning a 'Test' user
func AuthenticateSessionCookie(c *http.Cookie) (models.User, error) {
	return models.User{Username: "Test", Password: "abc"}, nil
}

package auth

import (
	"socialnetwork/models"
	"time"

	"github.com/google/uuid"
)

const (
	timeout    = 10 * time.Minute
	cookieName = "SessionID"
)

var cookieValue = GenerateNewUUID()

var SessionExpiration = time.Now().Add(timeout)

var sessionMap = map[string]*models.User{}

var reflectedSessionMap = map[*models.User]string{}

// generates a new UUID
func GenerateNewUUID() string {
	newUUID := uuid.New()
	// Convert the UUID to a string for display
	uuidString := newUUID.String()

	return uuidString
}

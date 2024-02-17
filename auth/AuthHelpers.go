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

var (
	cookieValue         = GenerateNewUUID()
	SessionExpiration   = time.Now().Add(timeout)
	sessionMap          = make(map[string]*models.User)
	reflectedSessionMap = make(map[*models.User]string)
)

// generates a new UUID
func GenerateNewUUID() string {
	newUUID := uuid.New()
	// Convert the UUID to a string for display
	uuidString := newUUID.String()
	return uuidString
}

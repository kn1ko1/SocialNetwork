package auth

import (
	"socialnetwork/models"
	"time"

	"github.com/google/uuid"
)

const (
	timeout    = 10 * time.Minute
	CookieName = "SessionID"
)

var (
	CookieValue = GenerateNewUUID()
	SessionMap  = make(map[string]*models.User)
	// followersMap = make(map[int][]int)
	// followingMap = make(map[int][]int)
	// groupsMap    = make(map[int][]int)
)

// generates a new UUID
func GenerateNewUUID() string {
	newUUID := uuid.New()
	// Convert the UUID to a string for display
	uuidString := newUUID.String()
	return uuidString
}

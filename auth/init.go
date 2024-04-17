package auth

var (
	DefaultManager ISessionManager
	CookieName     = "SessionID"
)

func init() {
	DefaultManager = NewSessionManager()
}

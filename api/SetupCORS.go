package api

import "net/http"

// May need tweaking to work properly

// The function sets up Cross-Origin Resource Sharing (CORS) headers for an HTTP response.
func SetupCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "HEAD, POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

	// Respond to preflight request with 200 OK status
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

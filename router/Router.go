package router

import (
	"net/http"
	"regexp"
)

type Router struct {
	Handlers map[*regexp.Regexp]http.Handler
}

func NewRouter() *Router {
	return &Router{Handlers: make(map[*regexp.Regexp]http.Handler)}
}
func (rt *Router) AddHandler(exp *regexp.Regexp, h http.Handler) {
	rt.Handlers[exp] = h
}
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// // Get Session Cookie
	// c, err := r.Cookie("Session")
	// if err != nil {
	// 	// Log Error
	// 	log.Println(err.Error())
	// 	// Return HTTP Status Unauthorized
	// 	//
	// 	// N.B. for simplicity of the example, we are simply returning
	// 	// an HTTP error. In the actual project, probably a JSON payload.
	// 	http.Error(w, "unauthorized", http.StatusUnauthorized)
	// 	return
	// }
	// // Authenticate Session Cookie - user variable discarded because user struct not used here...
	// _, err = auth.AuthenticateSessionCookie(c)
	// if err != nil {
	// 	// Same error as above - maker of request is unauthorized
	// 	log.Println(err.Error())
	// 	http.Error(w, "unauthorized", http.StatusUnauthorized)
	// 	return
	// }

	for exp, h := range rt.Handlers {
		if exp.MatchString(r.URL.Path) {
			h.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
	// w.WriteHeader(http.StatusNotFound)
	// enc := json.NewEncoder(w)
	// jsonPayload := transport.ErrorDTO{HTTPStatusCode: http.StatusNotFound, Message: "resource not found"}
	// err := enc.Encode(&jsonPayload)
	// if err != nil {
	// 	http.Error(w, "resource not found - additional error formulating response body", http.StatusNotFound)
	// }
}

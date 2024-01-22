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

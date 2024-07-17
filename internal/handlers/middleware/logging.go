package middleware

import (
	log "bank24/internal/logger"
	"net/http"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("new request from ", r.RemoteAddr, " as:", r.Method, " to: ", r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

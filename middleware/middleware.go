package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ForceContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Route='" + r.RequestURI + "' Method: '" + r.Method + "' ")
		next.ServeHTTP(w, r)
	})
}

func RegisterMiddleware(r *mux.Router) {
	log.Default().Println("Registering middleware")
	r.Use(LoggingMiddleware)
	r.Use(ForceContentTypeMiddleware)
	log.Default().Println("Middleware enabled.")
}

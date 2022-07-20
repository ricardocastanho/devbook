package middleware

import (
	"api/src/presenters"
	"api/src/support"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"\nMethod: %s - RequestURI: %s - Host: %s",
			r.Method,
			r.RequestURI,
			r.Host,
		)

		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := support.ValidateToken(r)

		if err != nil {
			presenters.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

package middleware

import (
	"api/src/presenters"
	"api/src/support"
	"net/http"
)

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

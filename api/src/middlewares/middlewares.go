package middlewares

import (
	"fmt"
	"net/http"

	"github.com/johnnyseubert/devbook/src/responses"
	"github.com/johnnyseubert/devbook/src/security"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.RequestURI)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

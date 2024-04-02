package app

import (
	"log"
	"net/http"

	"github.com/wandersonpaes/runners-api/internal/pkg/response"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
)

func logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidateToken(r); err != nil {
			response.ERR(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}
}

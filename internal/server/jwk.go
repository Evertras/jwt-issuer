package server

import (
	"net/http"

	"github.com/Evertras/jwt-issuer/internal/token"
)

func jwkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(token.JWKJson()))
	}
}

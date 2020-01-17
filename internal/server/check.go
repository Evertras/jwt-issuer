package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/Evertras/jwt-issuer/internal/token"
)

func checkHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		// Inefficient but who cares
		authorization = strings.TrimPrefix(authorization, "Bearer ")
		authorization = strings.TrimPrefix(authorization, "bearer ")
		authorization = strings.TrimSpace(authorization)

		claim, err := token.Parse(authorization)

		if err != nil {
			log.Println("Failed to parse token:", err)
			w.WriteHeader(401)
			return
		}

		log.Printf("Got claim: %+v", claim)
	}
}

package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Evertras/jwt-issuer/internal/token"
)

func checkHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		// Inefficient but who cares
		authorization = strings.TrimPrefix(authorization, "Bearer ")
		authorization = strings.TrimPrefix(authorization, "bearer ")
		authorization = strings.TrimSpace(authorization)

		claim, err := token.Parse(strings.NewReader(authorization))

		if err != nil {
			log.Println("Failed to parse token:", err)
			w.WriteHeader(401)
			return
		}

		json, err := claim.MarshalJSON()

		if err != nil {
			log.Println("Failed to marshal JSON:", err)
			w.WriteHeader(500)
			return
		}

		log.Printf("Got claim: %s", string(json))

		if claim.Expiration().Unix() < time.Now().Unix() {
			log.Println("Token expired")
			w.WriteHeader(401)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}
}

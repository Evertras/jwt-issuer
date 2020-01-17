package server

import (
	"log"
	"net/http"

	"github.com/Evertras/jwt-issuer/internal/token"
)

func generateHandler() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		user := req.Header.Get("X-User-ID")

		if user == "" {
			user = "fakeuser"
		}

		token, err := token.New(user)

		if err != nil {
			log.Println("Error making token:", err)
			w.WriteHeader(500)
			return
		}

		log.Printf("Created token for user ID %q", user)

		w.Write([]byte(token))
	}
}

package main

import (
	"log"

	"github.com/Evertras/jwt-issuer/internal/server"
)

func main() {
	s := server.New(":8080")

	log.Println("Running on :8080")

	log.Fatal(s.ListenAndServe())
}

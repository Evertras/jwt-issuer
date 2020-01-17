package server

import (
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
}

type server struct {
	httpServer *http.Server
}

func New(addr string) Server {
	router := http.NewServeMux()

	router.HandleFunc("/generate", generateHandler())

	return &server{
		httpServer: &http.Server{
			Addr:         addr,
			WriteTimeout: time.Second * 5,
			ReadTimeout: time.Second * 5,
			Handler: router,
		},
	}
}

func (s *server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

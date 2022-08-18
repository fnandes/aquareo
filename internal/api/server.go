package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

type Server struct {
	c    aquareo.Controller
	http *http.Server
}

func NewServer(config aquareo.Config, c aquareo.Controller) *Server {
	return &Server{c: c}
}

func (s *Server) Start(addr string) {
	log.Println("api: Server started")

	r := &handler{
		c: s.c,
	}

	s.http = &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(s.http.ListenAndServe())
}

func (s *Server) Stop(ctx context.Context) {
	if s.http != nil {
		log.Println("api: server stopped")
		_ = s.http.Shutdown(ctx)
	}
}

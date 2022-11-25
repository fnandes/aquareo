package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fnandes/aquareo/internal/aquareo"
)

type Server struct {
	ctrl aquareo.Controller
	http *http.Server
	cfg  aquareo.Config
}

func NewServer(ctrl aquareo.Controller, cfg aquareo.Config) *Server {
	return &Server{ctrl: ctrl, cfg: cfg}
}

func (s *Server) Start(addr string) {
	log.Println("api: Server started")

	r := &handler{
		ctrl: s.ctrl,
		cfg:  s.cfg,
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

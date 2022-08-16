package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

type Server struct {
	c aquareo.Controller
}

func NewServer(config aquareo.Config, c aquareo.Controller) *Server {
	return &Server{c}
}

func (s Server) Start(addr string) {
	log.Println("Starting API Server")

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	log.Fatal(http.ListenAndServe(addr, r))
}

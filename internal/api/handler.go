package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

type handler struct {
	ctrl aquareo.Controller
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()

	r.PathPrefix("/ui/").Handler(
		http.StripPrefix("/ui/", http.FileServer(http.Dir("ui/"))),
	)
	r.Use(corsMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		http.ServeFile(w, res, "ui/index.html")
	})
	r.HandleFunc("/api/config", h.GetConfig).Methods("GET")
	r.HandleFunc("/api/metrics/{bucket}", h.GetMetric).Methods("GET")

	r.ServeHTTP(w, req)
}

func (h *handler) GetMetric(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	items, err := h.ctrl.Storage().MetricStore(vars["bucket"]).List(40)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, _ := json.Marshal(items)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *handler) GetConfig(w http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(h.ctrl.Config())

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

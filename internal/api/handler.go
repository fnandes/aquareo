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
	cfg  aquareo.Config
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()

	r.PathPrefix("/ui/").Handler(
		http.StripPrefix("/ui/", http.FileServer(http.Dir("ui/"))),
	)
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		http.ServeFile(w, res, "ui/index.html")
	})
	r.HandleFunc("/api/config", h.GetConfig).Methods("GET")
	r.HandleFunc("/api/metrics/{bucket}", h.GetMetric).Methods("GET")
	r.HandleFunc("/api/metrics/{bucket}", h.AddMetricEntry).Methods("POST")

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

	json.NewEncoder(w).Encode(items)
}

func (h *handler) AddMetricEntry(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	var entry aquareo.MetricEntry
	if err := json.NewDecoder(req.Body).Decode(&entry); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	store := h.ctrl.Storage().MetricStore(vars["bucket"])
	if err := store.Put(entry.Timespan, entry.Value); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) GetConfig(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(h.ctrl.Config())
}

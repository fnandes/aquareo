package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fnandes/aquareo/internal/aquareo"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		http.ServeFile(w, res, "ui/index.html")
	})
	r.HandleFunc("/api/config", h.GetConfig).Methods("GET")
	r.HandleFunc("/api/metrics/{bucket}", h.GetMetric).Methods("GET")
	r.HandleFunc("/api/metrics/{bucket}", h.AddMetricEntry).Methods("POST")
	r.HandleFunc("/api/metrics/{bucket}/{timespan}", h.DeleteMetricEntry).Methods("DELETE")

	r.ServeHTTP(w, req)
}

func (h *handler) GetMetric(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	items, err := h.ctrl.Storage().MetricStore(vars["bucket"]).List(int(h.cfg.SystemSettings.MetricsLimit))
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

func (h *handler) DeleteMetricEntry(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	store := h.ctrl.Storage().MetricStore(vars["bucket"])

	key, err := strconv.ParseInt(vars["timespan"], 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := store.Delete(key); err != nil {
		log.Println(fmt.Errorf("DeleteMetricEntry: unable to delete: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetConfig(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(h.ctrl.Config())
}

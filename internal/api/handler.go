package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

type handler struct {
	c   aquareo.Controller
	cfg aquareo.Config
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("./ui")))

	r.HandleFunc("/", func(w http.ResponseWriter, res *http.Request) {
		http.ServeFile(w, res, "ui/index.html")
	})
	r.HandleFunc("/metrics", h.ListMetrics).Methods("GET")
	r.HandleFunc("/metrics/{key}", h.GetMetric).Methods("GET")

	r.ServeHTTP(w, req)
}

func (h *handler) GetMetric(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	items, err := h.c.Store().ReadAll(vars["key"], 40)
	if err != nil {
		log.Println("handler.GetMetric: Failed to get metric data: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, _ := json.Marshal(items)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

type metricResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *handler) ListMetrics(w http.ResponseWriter, req *http.Request) {
	var arr []metricResponse

	for _, s := range h.cfg.Sensors {
		arr = append(arr, metricResponse{
			Id:   s.Id,
			Name: s.Name,
		})
	}
	// add raspberry sensor
	arr = append(arr, metricResponse{
		Id:   aquareo.SensorSysTemp,
		Name: "Controller Temperature",
	})

	b, _ := json.Marshal(arr)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

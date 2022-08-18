package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

type handler struct {
	c aquareo.Controller
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r := mux.NewRouter()
	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("./web")))

	r.HandleFunc("/metric/{key}", h.GetMetric).Methods("GET")

	r.ServeHTTP(w, req)
}

func (h *handler) GetMetric(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	result, err := h.c.Store().ReadAll(vars["key"], 40)
	if err != nil {
		log.Println("handler.GetMetric: Failed to get metric data: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

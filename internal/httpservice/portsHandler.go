package httpservice

import (
	"encoding/json"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"github.com/ancarebeca/PortDomainService/internal/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type PortsHandler struct {
	Repository repository.Repository
}

// GetAllPorts returns all ports stored in the system */
func (h *PortsHandler) GetAllPorts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ports, err := h.Repository.GetAllPort()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ports)
}

// GetPort returns a por given an ID
func (h *PortsHandler) GetPort(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	w.WriteHeader(http.StatusOK)

	p, err := h.Repository.GetPort(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(p)
}

// AddOrUpdatesPort add a new port to the system. If it already exists it will be updated
func (h *PortsHandler) AddOrUpdatesPort(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := port.Port{}
	_ = json.NewDecoder(req.Body).Decode(&p)
	err := h.Repository.AddPort(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

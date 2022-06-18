package controller

import (
	"encoding/json"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type PortController struct {
	service.PortsService
}

// GetAllPorts returns all ports stored in the system */
func (c *PortController) GetAllPorts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ports, err := c.PortsService.GetAllPorts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ports); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetPort returns a por given an ID
func (c *PortController) GetPort(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	w.WriteHeader(http.StatusOK)

	p, err := c.PortsService.GetPort(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddOrUpdatesPort add a new domain to the system. If it already exists it will be updated
func (c *PortController) AddOrUpdatesPort(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := entity.Port{}
	_ = json.NewDecoder(req.Body).Decode(&p)
	err := c.PortsService.AddOrUpdatesPort(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

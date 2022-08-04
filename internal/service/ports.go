package service

import (
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
)

type Ports struct {
	Repository repository.Repository
}

type PortsService interface {
	GetAllPorts() ([]entity.Port, error)
	GetPort(ID string) (entity.Port, error)
	AddOrUpdatesPort(p entity.Port) error
}

// GetAllPorts returns all ports stored in the system */
func (h *Ports) GetAllPorts() ([]entity.Port, error) {
	return h.Repository.GetAllPorts()
}

// GetPort returns a por given an ID
func (h *Ports) GetPort(ID string) (entity.Port, error) {
	return h.Repository.GetPort(ID)
}

// AddOrUpdatesPort add a new domain to the system. If it already exists it will be updated
func (h *Ports) AddOrUpdatesPort(p entity.Port) error {
	return h.Repository.AddPort(p)
}

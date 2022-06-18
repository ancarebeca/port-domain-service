package repository

import (
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
)

type ReadWriter interface {
	ReadByID(id string) (entity.Port, error)
	Write(p entity.Port) error
	ReadAll() ([]entity.Port, error)
}

type Repository struct {
	DB ReadWriter
}

// GetPort by ID
func (r Repository) GetPort(ID string) (entity.Port, error) {
	p, err := r.DB.ReadByID(ID)
	if err != nil {
		return entity.Port{}, fmt.Errorf(fmt.Sprintf("Port %v was not found", ID))
	}
	return p, nil
}

// GetAllPorts in the system
func (r Repository) GetAllPorts() ([]entity.Port, error) {
	ports, err := r.DB.ReadAll()
	if err != nil {
		return ports, err
	}
	return ports, nil
}

// AddPort creates a new Port in the system
func (r Repository) AddPort(port entity.Port) error {
	err := r.DB.Write(port)
	if err != nil {
		return err
	}

	return nil
}

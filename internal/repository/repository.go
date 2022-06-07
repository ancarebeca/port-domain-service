package repository

import (
	"errors"
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"os"
)

type ReadWriter interface {
	ReadByID(id string) (port.Port, error)
	Write(p port.Port) error
	ReadAll() ([]port.Port, error)
}

type ObjectReader interface {
	Read(file *os.File) ([]port.Port, error)
}

type Repository struct {
	DB     ReadWriter
	Reader ObjectReader
}

func (r Repository) GetPort(id string) (port.Port, error) {

	p, err := r.DB.ReadByID(id)
	if err != nil {
		return port.Port{}, errors.New(fmt.Sprintf("Port %v was not found", id))
	}
	return p, nil
}

func (r Repository) GetAllPort() ([]port.Port, error) {
	ports, err := r.DB.ReadAll()
	if err != nil {
		return ports, err
	}
	return ports, nil
}

func (r Repository) AddPort(port port.Port) error {
	err := r.DB.Write(port)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) LoadPortsFromFile(file *os.File) error {
	ports, err := r.Reader.Read(file)
	if err != nil {
		return err
	}

	for _, p := range ports {
		err := r.DB.Write(p)
		if err != nil {
			return err
		}
	}
	return nil
}

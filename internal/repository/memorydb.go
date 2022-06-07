package repository

import (
	"errors"
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/port"
)

type MemDB struct {
	DB map[string]port.Port
}

func (m MemDB) ReadByID(id string) (port.Port, error) {
	p, ok := m.DB[id]
	if !ok {
		return port.Port{}, errors.New(fmt.Sprintf("Port %v was not found", id))
	}
	return p, nil
}

func (m MemDB) ReadAll() ([]port.Port, error) {
	var ports []port.Port

	for _, p := range m.DB {
		ports = append(ports, p)
	}

	return ports, nil
}

func (m MemDB) Write(p port.Port) error {
	m.DB[p.ID] = p
	return nil
}

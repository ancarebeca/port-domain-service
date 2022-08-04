package adapter

import (
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
)

type MemDB struct {
	DB map[string]entity.Port
}

func (m MemDB) ReadByID(id string) (entity.Port, error) {
	p, ok := m.DB[id]
	if !ok {
		return entity.Port{}, fmt.Errorf(fmt.Sprintf("Port %v was not found", id))
	}
	return p, nil
}

func (m MemDB) ReadAll() ([]entity.Port, error) {
	var ports []entity.Port

	for _, p := range m.DB {
		ports = append(ports, p)
	}

	return ports, nil
}

func (m MemDB) Write(p entity.Port) error {
	m.DB[p.ID] = p
	return nil
}

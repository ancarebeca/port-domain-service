package service

import (
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
	"io"
)

type ObjectReader interface {
	Read(r io.Reader) ([]entity.Port, error)
}

type Loader struct {
	ObjectReader
	repository.Repository
}

type LoaderService interface {
	Load(r io.Reader) error
}

// Load ports information from io.Reader into the system
func (l Loader) Load(r io.Reader) error {
	ports, err := l.Read(r)
	if err != nil {
		return err
	}

	for _, p := range ports {
		err := l.DB.Write(p)
		if err != nil {
			return err
		}
	}
	return nil
}

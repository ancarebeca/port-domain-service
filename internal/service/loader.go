package service

import (
	"encoding/json"
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
	"io"
)

type Loader struct {
	repository.Repository
}

type LoaderService interface {
	Load(r io.Reader) error
}

// Load into de system a json file that contains domain information
func (l Loader) Load(r io.Reader) error {
	dec := json.NewDecoder(r)

	_, err := dec.Token()
	if err != nil {
		return err
	}

	var token json.Token
	for dec.More() {
		token, err = dec.Token()
		if err != nil {
			return err
		}

		var p entity.Port
		err := dec.Decode(&p)
		if err != nil {
			return err
		}

		p.ID = fmt.Sprint(token)
		dbError := l.AddPort(p)
		if dbError != nil {
			return dbError
		}
	}

	return nil
}

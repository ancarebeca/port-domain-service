package reader

import (
	"encoding/json"
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"io"
)

type JSONReader struct {
}

// Read reads json object that contains domain information
func (jr JSONReader) Read(r io.Reader) ([]entity.Port, error) {
	var ports []entity.Port
	dec := json.NewDecoder(r)

	_, err := dec.Token()
	if err != nil {
		return ports, err
	}

	var token json.Token
	for dec.More() {
		token, err = dec.Token()
		if err != nil {
			return ports, err
		}

		var p entity.Port
		err := dec.Decode(&p)
		if err != nil {
			return ports, err
		}

		p.ID = fmt.Sprint(token)
		ports = append(ports, p)
	}

	return ports, nil
}

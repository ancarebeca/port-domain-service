package reader

import (
	"encoding/json"
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"os"
)

type Reader struct{}

// Read reads json file that contains port information
func (r Reader) Read(file *os.File) ([]port.Port, error) {
	var ports []port.Port
	dec := json.NewDecoder(file)

	// read open bracket
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

		var p port.Port
		err := dec.Decode(&p)
		if err != nil {
			return ports, err
		}

		p.ID = fmt.Sprint(token)
		ports = append(ports, p)
	}

	return ports, nil
}

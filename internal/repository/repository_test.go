package repository

import (
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"github.com/ancarebeca/PortDomainService/internal/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// TODO: missing test for unhappy path

func TestRepository(t *testing.T) {
	t.Run("loads ports from file", func(t *testing.T) {
		file := loadFixture(t)

		db := make(map[string]port.Port)
		memDB := MemDB{db}
		repository := Repository{
			DB:     memDB,
			Reader: reader.Reader{},
		}

		err := repository.LoadPortsFromFile(file)
		require.NoError(t, err)
		assert.Equal(t, expectedPorts["AEAJM"], db["AEAJM"])
		assert.Equal(t, expectedPorts["AEAUH"], db["AEAUH"])
	})

	t.Run("adds a new port", func(t *testing.T) {
		db := make(map[string]port.Port)
		memDB := MemDB{db}
		repository := Repository{
			DB:     memDB,
			Reader: reader.Reader{},
		}

		err := repository.AddPort(newPort)
		require.NoError(t, err)
		assert.Equal(t, newPort, db["BBBBB"])
	})

	t.Run("updates a port", func(t *testing.T) {
		db := make(map[string]port.Port)
		memDB := MemDB{db}
		repository := Repository{
			DB:     memDB,
			Reader: reader.Reader{},
		}

		err := repository.AddPort(updatedPort)
		require.NoError(t, err)
		assert.Equal(t, updatedPort, db["BBBBB"])
	})
}

func loadFixture(t *testing.T) *os.File {
	name := "ports.json"
	file, err := os.Open(fmt.Sprintf("fixtures/%s", name))
	require.NoError(t, err)
	return file
}

var newPort = port.Port{
	ID:          "BBBBB",
	Name:        "Ajman",
	City:        "Ajman",
	Country:     "United Arab Emirates",
	Alias:       []string{},
	Regions:     []string{},
	Coordinates: []float64{55.5136433, 25.4052165},
	Province:    "Ajman",
	Timezone:    "Asia/Dubai",
	Unlocs:      []string{"AEAJM"},
	Code:        "52000",
}

var updatedPort = port.Port{
	ID:          "BBBBB",
	Name:        "Ajman",
	City:        "Ajman",
	Country:     "Spain",
	Alias:       []string{},
	Regions:     []string{},
	Coordinates: []float64{55.5136433, 25.4052165},
	Province:    "Ajman",
	Timezone:    "Asia/Dubai",
	Unlocs:      []string{"AEAJM"},
	Code:        "52000",
}

var expectedPorts = map[string]port.Port{
	"AEAJM": port.Port{
		ID:          "AEAJM",
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	},
	"AEAUH": port.Port{
		ID:          "AEAUH",
		Name:        "Abu Dhabi",
		Coordinates: []float64{54.37, 24.47},
		City:        "Abu Dhabi",
		Province:    "Abu Z¸aby [Abu Dhabi]",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAUH"},
		Code:        "52001",
	},
}

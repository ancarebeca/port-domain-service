package service

import (
	"fmt"
	"github.com/ancarebeca/PortDomainService/internal/adapter"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
	"github.com/ancarebeca/PortDomainService/internal/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// TODO: missing test for unhappy path

func TestPortLoader(t *testing.T) {
	t.Run("loads ports from JSON", func(t *testing.T) {
		file := loadFixture(t)

		db := make(map[string]entity.Port)
		loader := Loader{
			reader.JSONReader{},
			repository.Repository{
				DB: adapter.MemDB{DB: db},
			},
		}

		err := loader.Load(file)
		require.NoError(t, err)
		assert.Equal(t, expectedPorts["AEAJM"], db["AEAJM"])
		assert.Equal(t, expectedPorts["AEAUH"], db["AEAUH"])
	})
}

func loadFixture(t *testing.T) *os.File {
	name := "ports.json"
	file, err := os.Open(fmt.Sprintf("fixture/%s", name))
	require.NoError(t, err)
	return file
}

var expectedPorts = map[string]entity.Port{
	"AEAJM": entity.Port{
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
	"AEAUH": entity.Port{
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

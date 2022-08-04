package repository

import (
	"github.com/ancarebeca/PortDomainService/internal/adapter"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO: missing test for unhappy path

func TestRepository(t *testing.T) {

	t.Run("adds a new domain", func(t *testing.T) {
		db := make(map[string]entity.Port)
		memDB := adapter.MemDB{DB: db}
		repository := Repository{
			DB: memDB,
		}

		err := repository.AddPort(newPort)
		require.NoError(t, err)
		assert.Equal(t, newPort, db["BBBBB"])
	})

	t.Run("updates a domain", func(t *testing.T) {
		db := make(map[string]entity.Port)
		memDB := adapter.MemDB{DB: db}
		repository := Repository{
			DB: memDB,
		}

		err := repository.AddPort(updatedPort)
		require.NoError(t, err)
		assert.Equal(t, updatedPort, db["BBBBB"])
	})
}

var newPort = entity.Port{
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

var updatedPort = entity.Port{
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

package controller

import (
	"bytes"
	"github.com/ancarebeca/PortDomainService/internal/adapter"
	"github.com/ancarebeca/PortDomainService/internal/domain/entity"
	"github.com/ancarebeca/PortDomainService/internal/domain/repository"
	"github.com/ancarebeca/PortDomainService/internal/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: missing test for unhappy path

func TestPortsController(t *testing.T) {

	s := service.Ports{
		Repository: repository.Repository{
			DB: adapter.MemDB{DB: loadDB()},
		},
	}

	c := PortController{
		&s,
	}

	t.Run("gets all ports", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/ports", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(c.GetAllPorts)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := []byte(`[{"id":"AEAJM","name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}]`)
		assert.JSONEq(t, string(expected), rr.Body.String())
	})

	t.Run("gets a domain by ID", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/ports/AEAJM", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		vars := map[string]string{
			"id": "AEAJM",
		}
		req = mux.SetURLVars(req, vars)

		handler := http.HandlerFunc(c.GetPort)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := []byte(`{"id":"AEAJM","name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`)
		assert.JSONEq(t, string(expected), rr.Body.String())
	})

	t.Run("creates a new domain", func(t *testing.T) {

		var jsonStr = []byte(`{"id":"FFFF","name":"Spain","city":"Spain","country":"Spain","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["FFFF"],"code":"52000"}`)
		req, err := http.NewRequest("POST", "/ports", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(c.AddOrUpdatesPort)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, string(jsonStr), rr.Body.String())
	})

	t.Run("updates the domain information", func(t *testing.T) {
		var jsonStr = []byte(`{"id":"AEAJM","name":"NEW NAME","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`)
		req, err := http.NewRequest("POST", "/ports", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(c.AddOrUpdatesPort)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, string(jsonStr), rr.Body.String())
	})
}

func loadDB() map[string]entity.Port {
	db := make(map[string]entity.Port)

	db["AEAJM"] = entity.Port{
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
	}
	return db
}

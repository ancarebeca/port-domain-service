package httpservice

import (
	"bytes"
	"github.com/ancarebeca/PortDomainService/internal/port"
	"github.com/ancarebeca/PortDomainService/internal/repository"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: missing test for unhappy path

func TestPortsHandler(t *testing.T) {

	portsHandler := PortsHandler{
		Repository: repository.Repository{
			DB: repository.MemDB{DB: loadDB()},
		},
	}

	t.Run("gets all ports", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/ports", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(portsHandler.GetAllPorts)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := []byte(`[{"id":"AEAJM","name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}]`)
		assert.JSONEq(t, string(expected), rr.Body.String())
	})

	t.Run("gets a port by ID", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/ports/AEAJM", nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		vars := map[string]string{
			"id": "AEAJM",
		}
		req = mux.SetURLVars(req, vars)

		handler := http.HandlerFunc(portsHandler.GetPort)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		expected := []byte(`{"id":"AEAJM","name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`)
		assert.JSONEq(t, string(expected), rr.Body.String())
	})

	t.Run("creates a new port", func(t *testing.T) {

		var jsonStr = []byte(`{"id":"FFFF","name":"Spain","city":"Spain","country":"Spain","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["FFFF"],"code":"52000"}`)
		req, err := http.NewRequest("POST", "/ports", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(portsHandler.AddOrUpdatesPort)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, string(jsonStr), rr.Body.String())
	})

	t.Run("updates the port information", func(t *testing.T) {
		var jsonStr = []byte(`{"id":"AEAJM","name":"NEW NAME","city":"Ajman","country":"United Arab Emirates","alias":[],"regions":[],"coordinates":[55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"}`)
		req, err := http.NewRequest("POST", "/ports", bytes.NewBuffer(jsonStr))
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(portsHandler.AddOrUpdatesPort)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.JSONEq(t, string(jsonStr), rr.Body.String())
	})
}

func loadDB() map[string]port.Port {
	db := make(map[string]port.Port)

	db["AEAJM"] = port.Port{
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

package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPersonCRUD(t *testing.T) {
	r := setup()

	req := httptest.NewRequest("POST", "/groups", bytes.NewBuffer([]byte(`{"name":"Group P"}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var g map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &g)
	groupID := g["id"].(string)

	person := []byte(`{
		"first_name":"Ivan",
		"last_name":"Ivanov",
		"birth_year":2000,
		"group_id":"` + groupID + `"
	}`)

	req = httptest.NewRequest("POST", "/people", bytes.NewBuffer(person))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, 201, w.Code)

	var p map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &p)

	id := p["id"].(string)
	require.NotEmpty(t, id)

	req = httptest.NewRequest("GET", "/people", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("GET", "/people/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	update := []byte(`{
		"first_name":"Petr",
		"last_name":"Petrov",
		"birth_year":1999,
		"group_id":"` + groupID + `"
	}`)

	req = httptest.NewRequest("PUT", "/people/"+id, bytes.NewBuffer(update))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("DELETE", "/people/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 204, w.Code)
}

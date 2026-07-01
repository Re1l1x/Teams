package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupCRUD(t *testing.T) {
	r := setup()

	createBody := []byte(`{"name":"Group A"}`)

	req := httptest.NewRequest("POST", "/groups", bytes.NewBuffer(createBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, 201, w.Code)

	var group map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &group)

	id := group["id"].(string)
	require.NotEmpty(t, id)

	req = httptest.NewRequest("GET", "/groups", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("GET", "/groups/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	updateBody := []byte(`{"name":"Group Updated"}`)

	req = httptest.NewRequest("PUT", "/groups/"+id, bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("DELETE", "/groups/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 204, w.Code)
}

func TestGroupPeopleEndpoints(t *testing.T) {
	r := setup()

	req := httptest.NewRequest("POST", "/groups", bytes.NewBuffer([]byte(`{"name":"Root"}`)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var g map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &g)

	id := g["id"].(string)

	req = httptest.NewRequest("GET", "/groups/"+id+"/people", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("GET", "/groups/"+id+"/people/all", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("GET", "/groups/"+id+"/count", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	req = httptest.NewRequest("GET", "/groups/"+id+"/count/all", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)
}

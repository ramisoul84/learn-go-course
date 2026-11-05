package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestApp returns the routes, with a fresh in-memory database.
func newTestApp(t *testing.T) http.Handler {
	t.Helper()
	store, err := OpenStore(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { store.Close() })
	app := &App{store: store}
	return app.routes()
}

// send runs one request, and returns the recorded response.
func send(
	h http.Handler, method, path, body string,
) *httptest.ResponseRecorder {
	in := strings.NewReader(body)
	req := httptest.NewRequest(method, path, in)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestCreateTodo(t *testing.T) {
	h := newTestApp(t)

	rec := send(h, "POST", "/todos", `{"title":"Milk"}`)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201", rec.Code)
	}
	var got Todo
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != 1 || got.Title != "Milk" {
		t.Errorf("got %+v", got)
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   int
	}{
		{"empty title", "POST", "/todos", `{"title":""}`, 400},
		{"broken JSON", "POST", "/todos", `oops`, 400},
		{"unknown ID", "GET", "/todos/99", "", 404},
		{"delete unknown", "DELETE", "/todos/99", "", 404},
		{"bad filter", "GET", "/todos?done=maybe", "", 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestApp(t)
			rec := send(h, tt.method, tt.path, tt.body)
			if rec.Code != tt.want {
				t.Errorf("status = %d, want %d", rec.Code, tt.want)
			}
		})
	}
}

func TestDoneFilter(t *testing.T) {
	h := newTestApp(t)
	send(h, "POST", "/todos", `{"title":"Milk","done":true}`)
	send(h, "POST", "/todos", `{"title":"Eggs"}`)

	rec := send(h, "GET", "/todos?done=true", "")

	var got []Todo
	err := json.NewDecoder(rec.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Title != "Milk" {
		t.Errorf("got %+v", got)
	}
}

// A JSON REST API for a to-do list, saved in a SQLite database.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Todo is one item. The tags set the names used in JSON.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// App holds what the handlers need. It only knows the interface.
type App struct {
	store Store
}

// writeJSON sends any value as a JSON response.
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// writeError sends an error as JSON.
func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// storeError answers when the store returns an error.
func storeError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	log.Println(err)
	writeError(w, http.StatusInternalServerError, "server error")
}

// readTodo decodes and checks the request body.
// If it returns false, the error is already sent.
func readTodo(w http.ResponseWriter, r *http.Request) (Todo, bool) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return todo, false
	}
	if strings.TrimSpace(todo.Title) == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return todo, false
	}
	return todo, true
}

// pathID reads the ID from the path. IDs start at 1, so a
// bad ID becomes 0, and the store answers "not found".
func pathID(r *http.Request) int {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return 0
	}
	return id
}

// list answers GET /todos.
func (a *App) list(w http.ResponseWriter, r *http.Request) {
	todos, err := a.store.All()
	if err != nil {
		storeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, todos)
}

// create answers POST /todos.
func (a *App) create(w http.ResponseWriter, r *http.Request) {
	todo, ok := readTodo(w, r)
	if !ok {
		return
	}
	todo, err := a.store.Add(todo)
	if err != nil {
		storeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, todo)
}

// get answers GET /todos/{id}.
func (a *App) get(w http.ResponseWriter, r *http.Request) {
	todo, err := a.store.Get(pathID(r))
	if err != nil {
		storeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

// update answers PUT /todos/{id}.
func (a *App) update(w http.ResponseWriter, r *http.Request) {
	todo, ok := readTodo(w, r)
	if !ok {
		return
	}
	todo.ID = pathID(r)
	if err := a.store.Update(todo); err != nil {
		storeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

// remove answers DELETE /todos/{id}.
func (a *App) remove(w http.ResponseWriter, r *http.Request) {
	if err := a.store.Delete(pathID(r)); err != nil {
		storeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	store, err := OpenStore("todos.db")
	if err != nil {
		log.Fatal(err)
	}
	app := &App{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", app.list)
	mux.HandleFunc("POST /todos", app.create)
	mux.HandleFunc("GET /todos/{id}", app.get)
	mux.HandleFunc("PUT /todos/{id}", app.update)
	mux.HandleFunc("DELETE /todos/{id}", app.remove)

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

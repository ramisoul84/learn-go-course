// Solution to the episode 22 exercise: a to-do must have a title.
// Run it with: go run ./solution
// Then try: curl -i -X POST localhost:8080/todos -d '{"title":""}'
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"
)

// Todo is one item. The tags set the names used in JSON.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// App keeps the to-dos in memory. Requests run at the same
// time, so a mutex protects them.
type App struct {
	mu     sync.Mutex
	todos  []Todo
	nextID int
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

// find returns the position of the to-do in the path, or -1.
// Call it with the mutex locked.
func (a *App) find(r *http.Request) int {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return -1
	}
	return slices.IndexFunc(a.todos, func(t Todo) bool {
		return t.ID == id
	})
}

// list answers GET /todos.
func (a *App) list(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	writeJSON(w, http.StatusOK, a.todos)
}

// create answers POST /todos.
func (a *App) create(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if strings.TrimSpace(todo.Title) == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.nextID++
	todo.ID = a.nextID
	a.todos = append(a.todos, todo)
	writeJSON(w, http.StatusCreated, todo)
}

// get answers GET /todos/{id}.
func (a *App) get(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	i := a.find(r)
	if i == -1 {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	writeJSON(w, http.StatusOK, a.todos[i])
}

// update answers PUT /todos/{id}.
func (a *App) update(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if strings.TrimSpace(todo.Title) == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	i := a.find(r)
	if i == -1 {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	todo.ID = a.todos[i].ID
	a.todos[i] = todo
	writeJSON(w, http.StatusOK, todo)
}

// remove answers DELETE /todos/{id}.
func (a *App) remove(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	i := a.find(r)
	if i == -1 {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}
	a.todos = slices.Delete(a.todos, i, i+1)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	app := &App{todos: []Todo{}}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", app.list)
	mux.HandleFunc("POST /todos", app.create)
	mux.HandleFunc("GET /todos/{id}", app.get)
	mux.HandleFunc("PUT /todos/{id}", app.update)
	mux.HandleFunc("DELETE /todos/{id}", app.remove)

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

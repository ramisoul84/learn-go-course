package main

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite" // registers the "sqlite" driver
)

// ErrNotFound means there is no to-do with that ID.
var ErrNotFound = errors.New("todo not found")

// Store is everything the handlers need from storage.
type Store interface {
	All() ([]Todo, error)
	Filter(done bool) ([]Todo, error)
	Add(t Todo) (Todo, error)
	Get(id int) (Todo, error)
	Update(t Todo) error
	Delete(id int) error
}

// SQLiteStore keeps the to-dos in a SQLite file.
type SQLiteStore struct {
	db *sql.DB
}

const schema = `CREATE TABLE IF NOT EXISTS todos (
	id    INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	done  BOOLEAN NOT NULL
)`

// OpenStore opens the database, and creates the table if needed.
func OpenStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// SQLite allows one writer at a time, so use one connection.
	db.SetMaxOpenConns(1)
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}
	return &SQLiteStore{db: db}, nil
}

// All returns every to-do.
func (s *SQLiteStore) All() ([]Todo, error) {
	rows, err := s.db.Query(
		"SELECT id, title, done FROM todos ORDER BY id")
	if err != nil {
		return nil, err
	}
	return collect(rows)
}

// Filter returns the to-dos that are done, or not done.
func (s *SQLiteStore) Filter(done bool) ([]Todo, error) {
	rows, err := s.db.Query(
		"SELECT id, title, done FROM todos WHERE done = ? ORDER BY id",
		done)
	if err != nil {
		return nil, err
	}
	return collect(rows)
}

// collect turns rows into to-dos.
func collect(rows *sql.Rows) ([]Todo, error) {
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

// Add saves a new to-do, and returns it with its ID.
func (s *SQLiteStore) Add(t Todo) (Todo, error) {
	res, err := s.db.Exec(
		"INSERT INTO todos (title, done) VALUES (?, ?)",
		t.Title, t.Done)
	if err != nil {
		return Todo{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	t.ID = int(id)
	return t, nil
}

// Get returns one to-do.
func (s *SQLiteStore) Get(id int) (Todo, error) {
	var t Todo
	row := s.db.QueryRow(
		"SELECT id, title, done FROM todos WHERE id = ?", id)
	err := row.Scan(&t.ID, &t.Title, &t.Done)
	if errors.Is(err, sql.ErrNoRows) {
		return Todo{}, ErrNotFound
	}
	return t, err
}

// Update replaces a to-do.
func (s *SQLiteStore) Update(t Todo) error {
	res, err := s.db.Exec(
		"UPDATE todos SET title = ?, done = ? WHERE id = ?",
		t.Title, t.Done, t.ID)
	if err != nil {
		return err
	}
	return mustChange(res)
}

// Delete removes a to-do.
func (s *SQLiteStore) Delete(id int) error {
	res, err := s.db.Exec(
		"DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return mustChange(res)
}

// mustChange returns ErrNotFound if no row was changed.
func mustChange(res sql.Result) error {
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

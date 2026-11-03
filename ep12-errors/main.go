package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// readConfig adds context to the error from os.ReadFile, and wraps it with %w.
func readConfig() error {
	_, err := os.ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}
	return nil
}

// start wraps again, so the message reads like a trail of breadcrumbs.
func start() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("start: %w", err)
	}
	return nil
}

// ErrNotFound is a sentinel error: a fixed value callers can check for.
var ErrNotFound = errors.New("not found")

func findUser(id int) (string, error) {
	if id != 1 {
		return "", fmt.Errorf("user %d: %w", id, ErrNotFound)
	}
	return "Alex", nil
}

// ValidationError is a custom error type that carries extra information.
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return e.Field + " is invalid"
}

func main() {
	// The program from the start of the episode.
	err := start()
	fmt.Println("error:", err)
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("config file missing, using defaults")
	}

	// errors.New and fmt.Errorf
	fmt.Println(errors.New("disk full"))
	name := "config.json"
	fmt.Println(fmt.Errorf("cannot open %s", name))

	// Sentinel errors and errors.Is
	_, err = findUser(7)
	fmt.Println(err)
	fmt.Println(err == ErrNotFound)          // false: err is wrapped
	fmt.Println(errors.Is(err, ErrNotFound)) // true: Is looks inside

	// Custom error types and errors.As
	ve := &ValidationError{Field: "email"}
	err = fmt.Errorf("save: %w", ve)
	fmt.Println(err)
	var target *ValidationError
	if errors.As(err, &target) {
		fmt.Println("bad field:", target.Field)
	}
}

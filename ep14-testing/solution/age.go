// Package age holds the ParseAge function from episode 12, ready to test.
package age

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrNotNumber  = errors.New("not a number")
	ErrOutOfRange = errors.New("out of range")
)

// ParseAge turns text into an age from 0 to 150.
func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("age %q: %w", s, ErrNotNumber)
	}
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("age %d: %w", age, ErrOutOfRange)
	}
	return age, nil
}

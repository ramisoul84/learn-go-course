// Solution to the episode 12 exercise. Run it with: go run ./solution
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Two sentinel errors, so callers can tell the problems apart with errors.Is.
var (
	ErrNotNumber  = errors.New("not a number")
	ErrOutOfRange = errors.New("out of range")
)

// ParseAge turns text into an age from 0 to 150.
func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", s, ErrNotNumber)
	}
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("parse age %d: %w", age, ErrOutOfRange)
	}
	return age, nil
}

func main() {
	for _, input := range []string{"30", "abc", "200"} {
		age, err := ParseAge(input)
		switch {
		case errors.Is(err, ErrNotNumber):
			fmt.Println("Please type a number:", err)
		case errors.Is(err, ErrOutOfRange):
			fmt.Println("That age isn't realistic:", err)
		case err != nil:
			fmt.Println("error:", err)
		default:
			fmt.Println("Age:", age)
		}
	}
}

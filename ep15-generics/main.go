package main

import (
	"cmp"
	"fmt"
)

// Without generics, you need one function per type.
func LargestInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LargestFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Largest works for every ordered type. Go also has a built-in max function that does this.
func Largest[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Index returns the position of target, or -1. comparable allows == on T.
func Index[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// Stack is a generic type: a stack of any type of value.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the top item. ok is false when the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item, true
}

func main() {
	// The program from the start of the episode.
	fmt.Println(Largest(3, 9))
	fmt.Println(Largest(1.5, 2.5))
	fmt.Println(Largest("apple", "pear"))

	// Without generics.
	fmt.Println(LargestInt(3, 9), LargestFloat(1.5, 2.5))

	// Naming the type yourself.
	fmt.Println(Largest[float64](1, 2.5))

	// comparable
	fmt.Println(Index([]string{"go", "rust", "zig"}, "rust"))
	fmt.Println(Index([]int{1, 2, 3}, 9))

	// A generic type.
	var names Stack[string]
	names.Push("Alex")
	names.Push("Sam")
	fmt.Println(names.Pop())
	fmt.Println(names.Pop())
	fmt.Println(names.Pop())
}

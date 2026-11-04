// Solution to the episode 15 exercise. Run it with: go run ./solution
package main

import (
	"fmt"
	"strings"
)

// Filter returns the items for which keep returns true.
func Filter[T any](items []T, keep func(T) bool) []T {
	var result []T
	for _, item := range items {
		if keep(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	even := func(n int) bool { return n%2 == 0 }
	fmt.Println(Filter([]int{1, 2, 3, 4, 5, 6}, even))

	startsWithG := func(s string) bool { return strings.HasPrefix(s, "g") }
	fmt.Println(Filter([]string{"go", "rust", "gopher"}, startsWithG))
}

// Solution to the episode 6 exercise. Run it with: go run ./solution
package main

import (
	"fmt"
	"slices"
)

// unique returns a new slice without duplicates, keeping the original order.
func unique(nums []int) []int {
	result := make([]int, 0, len(nums)) // room for every number, so append never has to grow
	for _, n := range nums {
		if !slices.Contains(result, n) {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	fmt.Println(unique([]int{3, 1, 3, 2, 1})) // [3 1 2]
	fmt.Println(unique(nil))                  // []
}

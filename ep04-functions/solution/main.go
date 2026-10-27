// Solution to the episode 4 exercise. Run it with: go run ./solution
package main

import "fmt"

// minMax returns the smallest and the largest number.
// It expects at least one number.
func minMax(nums ...int) (int, int) {
	lo, hi := nums[0], nums[0]
	for _, n := range nums {
		if n < lo {
			lo = n
		}
		if n > hi {
			hi = n
		}
	}
	return lo, hi
}

func main() {
	lo, hi := minMax(3, 9, 1, 7)
	fmt.Println("min:", lo, "max:", hi)
}

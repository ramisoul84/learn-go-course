// Solution to the episode 9 exercise. Run it with: go run ./solution
package main

import "fmt"

// swap swaps the values that a and b point to.
func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	// Go can also do it in one line: *a, *b = *b, *a
}

// swapCopies can't work: a and b are copies of the caller's values.
// Swapping the copies changes nothing outside this function.
func swapCopies(a, b int) {
	a, b = b, a
	_, _ = a, b
}

func main() {
	x, y := 1, 2
	swapCopies(x, y)
	fmt.Println(x, y) // 1 2
	swap(&x, &y)
	fmt.Println(x, y) // 2 1
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	// The surprise from the start of the episode: slicing shares memory.
	a := []int{1, 2, 3}
	b := a[:2]
	fmt.Println("a =", a)
	b[0] = 99
	fmt.Println("b =", b)
	fmt.Println("a =", a) // a changed too: a and b share one array

	// slices.Clone makes an independent copy.
	c := slices.Clone(b)
	c[0] = 1
	fmt.Println(a, b, c)

	// Arrays: a fixed size that is part of the type.
	var scores [3]int
	scores[0] = 90
	scores[2] = 75
	fmt.Println(scores, len(scores)) // [90 0 75] 3

	// Slices can grow with append. Always assign the result back.
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	fmt.Println(nums, len(nums))
	fmt.Println(nums[1:3], nums[:2], nums[2:])
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// len and cap while appending. The exact capacities depend on the Go version.
	var s []int
	for i := range 5 {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}

	// make, copy and the slices package.
	m := make([]int, 3, 10)
	fmt.Println(m, len(m), cap(m)) // [0 0 0] 3 10

	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Println(dst, n)

	unsorted := []int{30, 10, 20}
	slices.Sort(unsorted)
	fmt.Println(unsorted, slices.Contains(unsorted, 20))
}

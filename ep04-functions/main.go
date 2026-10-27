package main

import (
	"errors"
	"fmt"
)

// area takes two float64 values and returns one.
func area(width, height float64) float64 {
	return width * height
}

// divide returns two values: the result, and an error if something went wrong.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// split has named results: half and rest start at 0, and a bare return returns them.
func split(total int) (half, rest int) {
	half = total / 2
	rest = total - half
	return
}

// sum is variadic: it accepts any number of ints.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// apply takes a function as a parameter.
func apply(x int, f func(int) int) int {
	return f(x)
}

func main() {
	// The program from the start of the episode.
	fmt.Printf("Area: %.2f\n", area(3, 4))

	result, err := divide(10, 4)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10 / 4 =", result)

	_, err = divide(1, 0)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Named results.
	fmt.Println(split(7))

	// Variadic with no arguments.
	fmt.Println(sum())

	// Functions are values.
	double := func(x int) int {
		return x * 2
	}
	fmt.Println(double(21))
	fmt.Println(apply(5, double))
}

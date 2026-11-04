// Solution to the episode 17 exercise. Run it with: go run ./solution
package main

import "fmt"

// generate sends the numbers 1 to 100.
func generate(out chan<- int) {
	for i := 1; i <= 100; i++ {
		out <- i
	}
	close(out)
}

// square receives numbers, and sends their squares.
func square(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go generate(numbers)
	go square(numbers, squares)

	total := 0
	for s := range squares {
		total += s
	}
	fmt.Println("total:", total)
}

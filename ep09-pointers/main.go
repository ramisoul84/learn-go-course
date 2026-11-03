package main

import "fmt"

// doubleCopy gets a copy of n, so the caller's variable never changes.
func doubleCopy(n int) {
	n *= 2
}

// double gets the address of the caller's variable, and changes it through the pointer.
func double(n *int) {
	*n *= 2
}

func main() {
	// The program from the start of the episode.
	x := 10
	fmt.Println("before:", x)
	double(&x)
	fmt.Println("after: ", x)

	// & gives an address, * follows it.
	y := 42
	p := &y
	fmt.Println(p)  // an address; different on every run
	fmt.Println(*p) // 42
	*p = 100
	fmt.Println(y) // 100: changed through p
	fmt.Printf("%T %T\n", y, p)

	// Go passes copies.
	z := 10
	doubleCopy(z)
	fmt.Println(z) // still 10
	double(&z)
	fmt.Println(z) // 20

	// nil pointers: check before you follow them.
	var empty *int
	fmt.Println(empty) // <nil>
	if empty != nil {
		fmt.Println(*empty)
	}
	q := new(int)
	fmt.Println(*q) // 0
}

package main

import "fmt"

// Size uses iota: Small = 0, Medium = 1, Large = 2.
type Size int

const (
	Small Size = iota
	Medium
	Large
)

const appName = "Learn Go"

func main() {
	// The program from the start of the episode.
	name := "Alex"
	age := 30
	height := 1.80
	admin := true

	fmt.Printf("Name:   %s\n", name)
	fmt.Printf("Age:    %d\n", age)
	fmt.Printf("Height: %.2f m\n", height)
	fmt.Printf("Admin:  %t\n", admin)

	// var: the long form, with and without a type.
	var city string = "Berlin"
	var score = 9.5
	fmt.Println(city, score)

	// Zero values: every variable starts with a sensible default.
	var count int
	var price float64
	var empty string
	var active bool
	fmt.Println(count, price, empty, active)
	fmt.Printf("%q\n", empty)

	// Constants and iota.
	fmt.Println(appName, Small, Medium, Large)

	// Type conversion: Go never converts for you.
	total := float64(age) + height
	whole := int(height) // the decimal part is cut off: 1
	fmt.Println(total, whole)
	fmt.Printf("Type: %T\n", height)
}
